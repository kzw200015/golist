package middlewares

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kzw200015/go-list/args"
	"github.com/kzw200015/go-list/assets"
)

func HandleStatics() gin.HandlerFunc {
	sFS, err := fs.Sub(assets.Get(), "frontend/dist")
	if err != nil {
		log.Panicln(err)
	}
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api/") {
			c.Next()
			return
		}

		if strings.HasPrefix(path, "/file/") {
			fileServer := http.StripPrefix("/file/", http.FileServer(http.Dir(args.GetSrcPath())))
			c.Header("Content-Disposition", "attachment")
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		fileServer := http.FileServer(http.FS(sFS))
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
