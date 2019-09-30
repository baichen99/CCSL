package middlewares

import (
	"ccsl/configs"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

type errorHandler func(context.Context, string)

type tokenExtractor func(context.Context) (string, error)

type jwtConfig struct {
	ValidationKeyGetter jwt.Keyfunc
	ContextKey          string
	ErrorHandler        errorHandler
	CredentialsOptional bool
	Extractor           tokenExtractor
	Debug               bool
	EnableAuthOnOptions bool
	SigningMethod       jwt.SigningMethod
}

type JwtMiddleware struct {
	jwtConfig jwtConfig
}

// SignJWTToken is used for sign a JWT Token for clients
func SignJWTToken(userID uuid.UUID, role string) (string, error) {
	tokenID := uuid.NewV4()
	payload := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"iss":  configs.Conf.App.Name,                                                          // Issuer
		"iat":  time.Now().Unix(),                                                              // Issued At
		"nbf":  time.Now().Unix(),                                                              // Not Before
		"jti":  tokenID,                                                                        // JWT Token ID
		"exp":  time.Now().Add(time.Hour * time.Duration(configs.Conf.JWT.ExpireHours)).Unix(), // Expiration Time
		"user": userID.String(),                                                                // Username
		"role": role,                                                                           // Role of User: Admin/User/etc...
	})
	key, _ := jwt.ParseECPrivateKeyFromPEM([]byte(configs.Conf.JWT.PrivateKey))
	token, err := payload.SignedString(key)
	return token, err
}

func onError(ctx context.Context, err string) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.Values().Set(errinfo, err)
}

func NewJwtMiddleware(cfg ...jwtConfig) *JwtMiddleware {
	var c jwtConfig
	if len(cfg) == 0 {
		c = jwtConfig{}
	} else {
		c = cfg[0]
	}
	if c.ContextKey == "" {
		c.ContextKey = "JWT"
	}
	if c.ErrorHandler == nil {
		c.ErrorHandler = onError
	}
	if c.Extractor == nil {
		c.Extractor = FromAuthHeader
	}
	return &JwtMiddleware{jwtConfig: c}
}

func (m *JwtMiddleware) Get(ctx context.Context) (user string, role string) {
	headers := ctx.Values().Get(m.jwtConfig.ContextKey)
	if headers == nil {
		ctx.StopExecution()
	} else {
		payload, _ := headers.(*jwt.Token)
		params, _ := payload.Claims.(jwt.MapClaims)
		user = params["user"].(string)
		role = params["role"].(string)
	}
	return
}

func (m *JwtMiddleware) Serve(ctx context.Context) {
	if err := m.CheckJWT(ctx); err != nil {
		ctx.StopExecution()
		return
	}
	ctx.Next()
}

func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}
	return authHeaderParts[1], nil
}

func FromParameter(param string) tokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

func FromFirst(extractors ...tokenExtractor) tokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

func (m *JwtMiddleware) CheckJWT(ctx context.Context) error {
	if !m.jwtConfig.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}
	token, err := m.jwtConfig.Extractor(ctx)
	if err != nil {
		m.logf(ctx, "Error extracting JWT: %v", err)
	} else {
		m.logf(ctx, "Token extracted: %s", token)
	}
	if err != nil {
		m.jwtConfig.ErrorHandler(ctx, err.Error())
		return fmt.Errorf("Error extracting token: %v", err)
	}
	if token == "" {
		if m.jwtConfig.CredentialsOptional {
			m.logf(ctx, "No credentials found (CredentialsOptional=true)")
			return nil
		}
		errorMsg := "TokenNotFound"
		m.jwtConfig.ErrorHandler(ctx, errorMsg)
		m.logf(ctx, "Error: No credentials found (CredentialsOptional=false)")
		return fmt.Errorf(errorMsg)
	}
	parsedToken, err := jwt.Parse(token, m.jwtConfig.ValidationKeyGetter)
	if err != nil {
		m.jwtConfig.ErrorHandler(ctx, "ExpiredToken")
		return fmt.Errorf("Error parsing token: %v", err)
	}
	if m.jwtConfig.SigningMethod != nil && m.jwtConfig.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.jwtConfig.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		m.logf(ctx, "Error validating token algorithm: %s", message)
		m.jwtConfig.ErrorHandler(ctx, errors.New(message).Error())
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}
	if !parsedToken.Valid {
		m.logf(ctx, "Token is invalid")
		m.jwtConfig.ErrorHandler(ctx, "InvalidToken")
		return fmt.Errorf("Token is invalid")
	}
	m.logf(ctx, "JWT: %v", parsedToken)
	ctx.Values().Set(m.jwtConfig.ContextKey, parsedToken)
	return nil
}

func (m *JwtMiddleware) logf(ctx context.Context, format string, args ...interface{}) {
	if m.jwtConfig.Debug {
		ctx.Application().Logger().Debugf("JWT Middleware: "+format, args...)
	}
}
