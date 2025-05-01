package routes

import (
	"hello-world/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.NoRoute(handlers.NotFoundHandler)
	router := app.Group("/api/v1")
	{
		router.GET("/hello", handlers.HelloHandler)
		router.GET("/date", handlers.DateHandler)
		router.GET("/health", handlers.HealthHandler)
	}

}
