package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckHealth(c *gin.Context) {
	db, ok := c.Get("db")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "database does not exist",
		})
		return
	}

	err := db.(*gorm.DB).Exec("SELECT 1").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "database error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
