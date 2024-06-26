package middlewares

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
)

func HandleReq(serverEventChan chan string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path

		method := ctx.Request.Method

		from := ctx.ClientIP()

		requesLog := fmt.Sprintf("[Request] %s: %s, from %s", method, url, from)

		serverEventChan <- requesLog

		ctx.Next()
	}
}
