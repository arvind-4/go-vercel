package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}
