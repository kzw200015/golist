package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/go-list/handlers"
	"github.com/kzw200015/go-list/middlewares"
)

var router *gin.Engine

func GetRouter() *gin.Engine {
	return router
}

func init() {
	router = gin.Default()
	router.Use(middlewares.HandleErrors(), cors.Default(), middlewares.HandleStatics())
	api := router.Group("/api")
	{
		api.GET("/list", handlers.ListPath())
		api.GET("/down", handlers.Down(router))
	}
}
