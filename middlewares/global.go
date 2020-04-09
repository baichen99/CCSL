package middlewares

import (
	"ccsl/configs"

	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/microcosm-cc/bluemonday"
	uuid "github.com/satori/go.uuid"
)

var (
	// JwtConf is config for JSON Web Token
	JwtConf = jwtConfig{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			key, err := jwt.ParseECPublicKeyFromPEM([]byte(configs.Conf.JWT.PublicKey))
			return key, err
		},
		SigningMethod: jwt.SigningMethodES512,
		Debug:         configs.Conf.JWT.Debug,
		ContextKey:    "JWT",
	}

	// CorsConf is config for Cross-Origin Resource Sharing
	CorsConf = cors.Options{
		AllowedOrigins:   configs.Conf.App.CORS,
		AllowCredentials: configs.Conf.App.AllowCookie,
		AllowedMethods:   []string{iris.MethodOptions, iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete},
	}

	// XSS filter, you may check out github.com/microcosm-cc/bluemonday/policies.go
	// UGCPolicy looks good for me.
	bluemondayConfig = bluemonday.UGCPolicy()
)

// CheckToken is a user authorization middleware, checks jwt token
var CheckToken = NewJwtMiddleware(JwtConf).Serve

// GetJWTParams returns JSON web token payload
var GetJWTParams = NewJwtMiddleware(JwtConf).Get

// CheckUserRole checks user role with multiple role values
func CheckUserRole(roles []string) func(ctx context.Context) {
	return NewRoleMiddleware(roleConfig{
		Roles: roles,
	}).Serve
}

// FilterUserInput filter XSS from users' input
func FilterUserInput(input string) (cleanData string) {
	cleanData = bluemondayConfig.Sanitize(input)
	return
}

// BeforeHandleRequest is a global middleware before handle a request
func BeforeHandleRequest(ctx iris.Context) {
	defer ctx.Next()
	requestID := uuid.NewV4().String()
	ctx.Values().Set("RequestID", requestID)
	ctx.Application().Logger().Infof("=====================REQUEST " + requestID + " START====================")
}

// AfterHandleRequest is a global middleware after handle a request
func AfterHandleRequest(ctx iris.Context) {
	defer ctx.Next()
	requestID := ctx.Values().Get("RequestID").(string)
	ctx.Application().Logger().Infof("=====================REQUEST " + requestID + " END======================")
}
