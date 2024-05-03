package middlewares

import "github.com/gin-gonic/gin"

// SSEHeader is a middleware that sets the necessary headers for Server-Sent Events (SSE).
// It sets the "Content-Type" header to "text/event-stream", the "Cache-Control" header to "no-cache",
// and the "Connection" header to "keep-alive".
func SSEHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		ctx.Next()
	}
}
