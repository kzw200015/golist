package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/go-list/handlers"
	"github.com/kzw200015/go-list/middlewares"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.HandleErrors()).Use(cors.Default()).Use(middlewares.HandleStatics())
	api := r.Group("/api")
	{
		api.GET("/list", handlers.ListPath)
	}

	return r
}
