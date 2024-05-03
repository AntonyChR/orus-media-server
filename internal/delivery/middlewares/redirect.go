package middlewares

import (
	"log"
	"net/http"
	"slices"
	"strings"

	gin "github.com/gin-gonic/gin"
)

// Frontend is an SPA, so if the user refreshes the page on a route that is not the root
// the server will return a 404 error. This middleware will redirect the user to the root
// where the frontend is served.
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
