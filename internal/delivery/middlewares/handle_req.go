package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleReq(serverEventChan chan string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path

		method := ctx.Request.Method

		from := ctx.ClientIP()

		requesLog := fmt.Sprintf("[%s] %s, from %s", method, url, from)

		go func() {
			serverEventChan <- requesLog
		}()

		ctx.Next()
	}
}
