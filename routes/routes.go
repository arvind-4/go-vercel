package routes

import (
	"hello-world/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.NoRoute(handlers.NotFoundHandler)
	router := app.Group("/api/v1")
	{
		router.GET("/hello", handlers.HelloHandler)
		router.GET("/health", handlers.HealthHandler)
	}

}
