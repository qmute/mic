package mic

import (
	"context"

	"github.com/gin-gonic/gin"
)

var activeCtxKey = "gin_trace_context"

// open trace middleware, 使用全局Tracer
func GinTraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := Trace(c, c.FullPath())
		c.Set(activeCtxKey, ctx)
		defer span.Finish()
		c.Next()
	}
}

// MustGetCtx extracts ctx（with trace span） from gin.Context. It panics if ctx was not set.
func MustGetCtx(c *gin.Context) context.Context {
	return c.MustGet(activeCtxKey).(context.Context)
}
