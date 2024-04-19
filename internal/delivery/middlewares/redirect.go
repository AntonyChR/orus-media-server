package middlewares

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectToRoot() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path
		if url == "/" {
			ctx.Next()
			return
		}

		splittedUrl := strings.Split(url, "/")

		s := []string{"movie", "movies", "series"}
		if slices.Contains(s, splittedUrl[1]) {
			ctx.Redirect(http.StatusMovedPermanently, "/")
			log.Printf("Redirect: \"%s\" -> \"%s\"", url, "\"/\"")
		}

		ctx.Next()
	}
}
