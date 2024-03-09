package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Run() {
	r := gin.Default()
	r.GET("/CheckHealth", checkHealth)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
