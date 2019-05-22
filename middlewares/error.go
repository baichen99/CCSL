package middlewares

import "github.com/kataras/iris"

const (
	message string = "message"
	errinfo string = "error"
)

const (
	// BadRequest 400 Error
	BadRequest string = "BadRequest"
	// Unauthorized 401 Error
	Unauthorized string = "Unauthorized"
	// Forbidden 403 Error
	Forbidden string = "Forbidden"
	// MethodError 405 Error
	MethodError string = "MethodError"
	// NotFound 404 Error
	NotFound string = "NotFound"
	// ConfictError 409 Error
	ConfictError string = "ConfictError"
	// UnprocessableEntity 422 Error
	UnprocessableEntity string = "UnprocessableEntity"
	// InternalError 500 Error
	InternalError string = "InternalError"
	// UnknownError other error code
	UnknownError string = "UnknownError"
)

// ErrorHandler handles request error
func ErrorHandler(ctx iris.Context) {
	errMessage := ctx.Values().GetString(errinfo)
	var msg string
	switch ctx.GetStatusCode() {
	case iris.StatusBadRequest: // 400 Error
		msg = ctx.Translate(BadRequest)
	case iris.StatusUnauthorized: //401 Error
		msg = ctx.Translate(Unauthorized)
	case iris.StatusForbidden: // 403 Error
		msg = ctx.Translate(Forbidden)
	case iris.StatusNotFound: // 404 Error
		msg = ctx.Translate(NotFound)
	case iris.StatusMethodNotAllowed: // 405 Error
		msg = ctx.Translate(MethodError)
	case iris.StatusConflict: // 409 Error
		msg = ctx.Translate(ConfictError)
	case iris.StatusUnprocessableEntity: // 422 Error
		msg = ctx.Translate(UnprocessableEntity)
	case iris.StatusInternalServerError: // 500 Error
		msg = ctx.Translate(InternalError)
	default:
		msg = ctx.Translate(UnknownError)
	}
	if errMessage != "" {
		ctx.JSON(iris.Map{
			message: msg,
			errinfo: ctx.Translate(errMessage),
		})
		return
	}
	ctx.JSON(iris.Map{
		message: msg,
	})
}
