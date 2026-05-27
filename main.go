package main

import (
	"hello-world/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.RegisterRoutes(app)
	app.Run()
}
