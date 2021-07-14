package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		e := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if e != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, e)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": e.Error(),
			})
		}
	}
}
