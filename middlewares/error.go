package middlewares

import (
	"github.com/kataras/iris/v12"
)

const (
	message             = "message"
	errinfo             = "error"
	BadRequest          = "BadRequest"          // BadRequest 400 Error
	Unauthorized        = "Unauthorized"        // Unauthorized 401 Error
	Forbidden           = "Forbidden"           // Forbidden 403 Error
	MethodError         = "MethodError"         // MethodError 405 Error
	NotFound            = "NotFound"            // NotFound 404 Error
	ConfictError        = "ConfictError"        // ConfictError 409 Error
	UnprocessableEntity = "UnprocessableEntity" // UnprocessableEntity 422 Error
	TooManyRequests     = "TooManyRequests"     // TooManyRequests 429 Error
	InternalError       = "InternalError"       // InternalError 500 Error
	URLNotFound         = "URLNotFound"         // URLNotFound 404 Error
	UnknownError        = "UnknownError"        // UnknownError other error code
)

// ErrorHandler handles request error
func ErrorHandler(ctx iris.Context) {
	errMessage := ctx.Values().GetString(errinfo)
	var msg string
	switch ctx.GetStatusCode() {
	case iris.StatusBadRequest: // 400 Error
		msg = BadRequest
	case iris.StatusUnauthorized: //401 Error
		msg = Unauthorized
	case iris.StatusForbidden: // 403 Error
		msg = Forbidden
	case iris.StatusNotFound: // 404 Error
		msg = NotFound
	case iris.StatusMethodNotAllowed: // 405 Error
		msg = MethodError
	case iris.StatusConflict: // 409 Error
		msg = ConfictError
	case iris.StatusUnprocessableEntity: // 422 Error
		msg = UnprocessableEntity
	case iris.StatusTooManyRequests: // 429 Error
		msg = TooManyRequests
	case iris.StatusInternalServerError: // 500 Error
		msg = InternalError
	default:
		msg = UnknownError
	}
	msg = ctx.Tr(msg)
	err := ctx.Tr(errMessage)
	if errMessage != "" {
		ctx.JSON(iris.Map{
			message: msg,
			errinfo: err,
		})
		return
	}
	ctx.JSON(iris.Map{
		message: ctx.Tr(URLNotFound),
		errinfo: nil,
	})
	ctx.StopExecution()
}
