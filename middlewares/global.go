package middlewares

import (
	"ccsl/configs"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
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
	CorsConf = corsConfig{
		AllowedOrigins:   configs.Conf.App.CORS,
		AllowCredentials: configs.Conf.App.AllowCookie,
	}
	// I18nConf is config for internationalization
	I18nConf = i18nConfig{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/en-US.ini",
			"zh-CN": "./locales/zh-CN.ini",
		},
	}
)

// CheckJWTToken is a user authorization middleware
var CheckJWTToken = NewJwtMiddleware(JwtConf).Serve

// GetJWTParams returns JSON web token payload
var GetJWTParams = NewJwtMiddleware(JwtConf).Get

//CheckAdmin checks if user is admin or super
var CheckAdmin = NewRoleMiddleware(roleConfig{
	Role: "admin",
}).Serve

//CheckSuper checks if user is super
var CheckSuper = NewRoleMiddleware(roleConfig{
	Role: "super",
}).Serve

// BeforeHandleRequest is a global middleware before handle a request
func BeforeHandleRequest(ctx iris.Context) {
	defer ctx.Next()
	ctx.Application().Logger().Infof("=====================REQUEST START====================")
}

// AfterHandleRequest is a global middleware after handle a request
func AfterHandleRequest(ctx iris.Context) {
	defer ctx.Next()
	ctx.Application().Logger().Infof("=====================REQUEST  END=====================")
}
