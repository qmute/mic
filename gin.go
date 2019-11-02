package mic

import (
	"context"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	regHash = regexp.MustCompile(`[0-9a-zA-Z]{17,}`)
	regCrop = regexp.MustCompile(`/\d{2,}x\d{2,}$`)
	regNum1 = regexp.MustCompile(`/\d+$`)
	regNum2 = regexp.MustCompile(`/\d+/`)
)

// 去除url中的常见变量，替换为占位符
func normalize(s string) string {
	s = strings.Split(s, "?")[0]
	s = regHash.ReplaceAllString(s, ":hash")
	s = regCrop.ReplaceAllString(s, "/:crop")
	s = regNum1.ReplaceAllString(s, "/:num")
	return regNum2.ReplaceAllString(s, "/:num/")
}

var activeCtxKey = "gin_trace_context"

// open trace middleware, 使用全局Tracer
func GinTraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := normalize(c.Request.URL.EscapedPath())
		ctx, span := Trace(c, path)
		c.Set(activeCtxKey, ctx)
		defer span.Finish()
		c.Next()
	}
}

// MustGetCtx extracts ctx（with trace span） from gin.Context. It panics if ctx was not set.
func MustGetCtx(c *gin.Context) context.Context {
	return c.MustGet(activeCtxKey).(context.Context)
}
