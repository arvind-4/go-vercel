package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DateHandler(c *gin.Context) {
	todayDate := time.Now().Format("2006-01-02")
	c.JSON(http.StatusOK, gin.H{
		"date": todayDate,
	})
}
