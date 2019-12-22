package middlewares

import (
	"ccsl/utils"
	"errors"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// LimitHandler returns a request limiter
func LimitHandler(l *limiter.Limiter) context.Handler {
	return func(ctx context.Context) {
		error := tollbooth.LimitByRequest(l, ctx.ResponseWriter(), ctx.Request())
		if error != nil {
			ctx.ContentType(l.GetMessageContentType())
			utils.LogInfo(ctx, "Request reached rate limit")
			utils.SetResponseError(ctx, iris.StatusTooManyRequests, "LimitMiddleware", errors.New("TryAgainLater"))
			return
		}
		ctx.Next()
	}
}

// CheckRateLimit checks request rate limit
func CheckRateLimit(max float64) context.Handler {
	limiter := tollbooth.NewLimiter(max, nil)
	return LimitHandler(limiter)
}
