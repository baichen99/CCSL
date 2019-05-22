package middlewares

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kataras/iris/context"
)

// corsConfig is a configuration container to setup the CORS middleware.
type corsConfig struct {
	AllowedOrigins     []string
	AllowOriginFunc    func(origin string) bool
	AllowedMethods     []string
	AllowedHeaders     []string
	ExposedHeaders     []string
	MaxAge             int
	AllowCredentials   bool
	OptionsPassthrough bool
	Debug              bool
}

// CorsMiddleware http handler
type CorsMiddleware struct {
	// Debug logger
	Log *log.Logger
	// Normalized list of plain allowed origins
	allowedOrigins []string
	// List of allowed origins containing wildcards
	allowedWOrigins []wildcard
	// Optional origin validator function
	allowOriginFunc func(origin string) bool
	// Normalized list of allowed headers
	allowedHeaders []string
	// Normalized list of allowed methods
	allowedMethods []string
	// Normalized list of exposed headers
	exposedHeaders []string
	maxAge         int
	// Set to true when allowed origins contains a "*"
	allowedOriginsAll bool
	// Set to true when allowed headers contains a "*"
	allowedHeadersAll bool
	allowCredentials  bool
	optionPassthrough bool
}

// NewCorsMiddleware creates a new Cors handler with the provided options.
func NewCorsMiddleware(options corsConfig) context.Handler {
	c := &CorsMiddleware{
		exposedHeaders:    convert(options.ExposedHeaders, http.CanonicalHeaderKey),
		allowOriginFunc:   options.AllowOriginFunc,
		allowCredentials:  options.AllowCredentials,
		maxAge:            options.MaxAge,
		optionPassthrough: options.OptionsPassthrough,
	}
	if options.Debug {
		c.Log = log.New(os.Stdout, "[cors] ", log.LstdFlags)
	}
	// Allowed Origins
	if len(options.AllowedOrigins) == 0 {
		if options.AllowOriginFunc == nil {
			// Default is all origins
			c.allowedOriginsAll = true
		}
	} else {
		c.allowedOrigins = []string{}
		c.allowedWOrigins = []wildcard{}
		for _, origin := range options.AllowedOrigins {
			// Normalize
			origin = strings.ToLower(origin)
			if origin == "*" {
				// If "*" is present in the list, turn the whole list into a match all
				c.allowedOriginsAll = true
				c.allowedOrigins = nil
				c.allowedWOrigins = nil
				break
			} else if i := strings.IndexByte(origin, '*'); i >= 0 {
				// Split the origin in two: start and end string without the *
				w := wildcard{origin[0:i], origin[i+1:]}
				c.allowedWOrigins = append(c.allowedWOrigins, w)
			} else {
				c.allowedOrigins = append(c.allowedOrigins, origin)
			}
		}
	}

	// Allowed Headers
	if len(options.AllowedHeaders) == 0 {
		// Use sensible defaults
		c.allowedHeaders = []string{"Origin", "Accept", "Content-Type", "X-Requested-With"}
	} else {
		// Origin is always appended as some browsers will always request for this header at preflight
		c.allowedHeaders = convert(append(options.AllowedHeaders, "Origin"), http.CanonicalHeaderKey)
		for _, h := range options.AllowedHeaders {
			if h == "*" {
				c.allowedHeadersAll = true
				c.allowedHeaders = nil
				break
			}
		}
	}
	// Allowed Methods
	if len(options.AllowedMethods) == 0 {
		// Default is spec's "simple" methods
		c.allowedMethods = []string{"GET", "POST", "HEAD"}
	} else {
		c.allowedMethods = convert(options.AllowedMethods, strings.ToUpper)
	}
	return c.Serve
}

// Default creates a new Cors handler with default options.
func Default() context.Handler {
	return NewCorsMiddleware(corsConfig{})
}

// AllowAll create a new Cors handler with permissive configuration allowing all
// origins with all standard methods with any header and credentials.
func AllowAll() context.Handler {
	return NewCorsMiddleware(corsConfig{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

// Serve apply the CORS specification on the request, and add relevant CORS headers
// as necessary.
func (c *CorsMiddleware) Serve(ctx context.Context) {
	if ctx.Method() == http.MethodOptions && ctx.GetHeader("Access-Control-Request-Method") != "" {
		c.logf(ctx, "Serve: Preflight request")
		c.handlePreflight(ctx)
		if c.optionPassthrough { // handle the options by routes.
			ctx.Next()
			return
		}
		if !ctx.IsStopped() {
			ctx.StatusCode(http.StatusOK)
			ctx.StopExecution()
		}
		return
	}
	c.logf(ctx, "Serve: Actual request")
	c.handleActualRequest(ctx)
	ctx.Next()
}

// handlePreflight handles pre-flight CORS requests
func (c *CorsMiddleware) handlePreflight(ctx context.Context) {
	origin := ctx.GetHeader("Origin")
	if ctx.Method() != http.MethodOptions {
		c.logf(ctx, "Preflight aborted: %s!=OPTIONS", ctx.Method())
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()
		return
	}
	// Always set Vary headers.
	ctx.Header("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
	if origin == "" {
		c.logf(ctx, "Preflight aborted: empty origin")
		return
	}
	if !c.isOriginAllowed(origin) {
		c.logf(ctx, "Preflight aborted: origin '%s' not allowed", origin)
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()
		return
	}

	reqMethod := ctx.GetHeader("Access-Control-Request-Method")
	if !c.isMethodAllowed(reqMethod) {
		c.logf(ctx, "Preflight aborted: method '%s' not allowed", reqMethod)
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()

		return
	}
	reqHeaders := parseHeaderList(ctx.GetHeader("Access-Control-Request-Headers"))
	if !c.areHeadersAllowed(reqHeaders) {
		c.logf(ctx, "Preflight aborted: headers '%v' not allowed", reqHeaders)
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()
		return
	}
	if c.allowedOriginsAll && !c.allowCredentials {
		ctx.Header("Access-Control-Allow-Origin", "*")
	} else {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	ctx.Header("Access-Control-Allow-Methods", strings.ToUpper(reqMethod))
	if len(reqHeaders) > 0 {
		ctx.Header("Access-Control-Allow-Headers", strings.Join(reqHeaders, ", "))
	}
	if c.allowCredentials {
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}
	if c.maxAge > 0 {
		ctx.Header("Access-Control-Max-Age", strconv.Itoa(c.maxAge))
	}
	c.logf(ctx, "Preflight response headers: %v", ctx.ResponseWriter().Header())
}

// handleActualRequest handles simple cross-origin requests, actual request or redirects
func (c *CorsMiddleware) handleActualRequest(ctx context.Context) {
	origin := ctx.GetHeader("Origin")

	if ctx.Method() == http.MethodOptions {
		c.logf(ctx, "Actual request no headers added: method == %s", ctx.Method())
		//
		ctx.StatusCode(http.StatusMethodNotAllowed)
		ctx.StopExecution()
		//
		return
	}
	// Always set Vary, see https://github.com/rs/cors/issues/10
	ctx.ResponseWriter().Header().Add("Vary", "Origin")
	if origin == "" {
		c.logf(ctx, "Actual request no headers added: missing origin")
		return
	}
	if !c.isOriginAllowed(origin) {
		c.logf(ctx, "Actual request no headers added: origin '%s' not allowed", origin)
		//
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()
		//
		return
	}

	if !c.isMethodAllowed(ctx.Method()) {
		c.logf(ctx, "Actual request no headers added: method '%s' not allowed", ctx.Method())
		ctx.StatusCode(http.StatusForbidden)
		ctx.StopExecution()
		return
	}
	if c.allowedOriginsAll && !c.allowCredentials {
		ctx.Header("Access-Control-Allow-Origin", "*")
	} else {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	if len(c.exposedHeaders) > 0 {
		ctx.Header("Access-Control-Expose-Headers", strings.Join(c.exposedHeaders, ", "))
	}
	if c.allowCredentials {
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}
	c.logf(ctx, "Actual response added headers: %v", ctx.ResponseWriter().Header())
}

// convenience method. checks if debugging is turned on before printing
func (c *CorsMiddleware) logf(ctx context.Context, format string, args ...interface{}) {
	ctx.Application().Logger().Debugf(format, args...)
}

// isOriginAllowed checks if a given origin is allowed to perform cross-domain requests
// on the endpoint
func (c *CorsMiddleware) isOriginAllowed(origin string) bool {
	if c.allowOriginFunc != nil {
		return c.allowOriginFunc(origin)
	}
	if c.allowedOriginsAll {
		return true
	}
	origin = strings.ToLower(origin)
	for _, o := range c.allowedOrigins {
		if o == origin {
			return true
		}
	}
	for _, w := range c.allowedWOrigins {
		if w.match(origin) {
			return true
		}
	}
	return false
}

// isMethodAllowed checks if a given method can be used as part of a cross-domain request
// on the endpoing
func (c *CorsMiddleware) isMethodAllowed(method string) bool {
	if len(c.allowedMethods) == 0 {
		// If no method allowed, always return false, even for preflight request
		return false
	}
	method = strings.ToUpper(method)
	if method == http.MethodOptions {
		// Always allow preflight requests
		return true
	}
	for _, m := range c.allowedMethods {
		if m == method {
			return true
		}
	}
	return false
}

// areHeadersAllowed checks if a given list of headers are allowed to used within
// a cross-domain request.
func (c *CorsMiddleware) areHeadersAllowed(requestedHeaders []string) bool {
	if c.allowedHeadersAll || len(requestedHeaders) == 0 {
		return true
	}
	for _, header := range requestedHeaders {
		header = http.CanonicalHeaderKey(header)
		found := false
		for _, h := range c.allowedHeaders {
			if h == header {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

const toLower = 'a' - 'A'

type converter func(string) string

type wildcard struct {
	prefix string
	suffix string
}

func (w wildcard) match(s string) bool {
	return len(s) >= len(w.prefix+w.suffix) && strings.HasPrefix(s, w.prefix) && strings.HasSuffix(s, w.suffix)
}

// convert converts a list of string using the passed converter function
func convert(s []string, c converter) []string {
	out := []string{}
	for _, i := range s {
		out = append(out, c(i))
	}
	return out
}

// parseHeaderList tokenize + normalize a string containing a list of headers
func parseHeaderList(headerList string) []string {
	l := len(headerList)
	h := make([]byte, 0, l)
	upper := true
	// Estimate the number headers in order to allocate the right splice size
	t := 0
	for i := 0; i < l; i++ {
		if headerList[i] == ',' {
			t++
		}
	}
	headers := make([]string, 0, t)
	for i := 0; i < l; i++ {
		b := headerList[i]
		if b >= 'a' && b <= 'z' {
			if upper {
				h = append(h, b-toLower)
			} else {
				h = append(h, b)
			}
		} else if b >= 'A' && b <= 'Z' {
			if !upper {
				h = append(h, b+toLower)
			} else {
				h = append(h, b)
			}
		} else if b == '-' || b == '_' || (b >= '0' && b <= '9') {
			h = append(h, b)
		}

		if b == ' ' || b == ',' || i == l-1 {
			if len(h) > 0 {
				// Flush the found header
				headers = append(headers, string(h))
				h = h[:0]
				upper = true
			}
		} else {
			upper = b == '-' || b == '_'
		}
	}
	return headers
}
