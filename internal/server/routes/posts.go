package routes

import (
	"net/http"

	dbApi "bikesense-web/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostTrip(c *gin.Context) {
	var trip dbApi.Trip

	if err := c.ShouldBindJSON(&trip); err != nil {
		example := dbApi.Trip{
			BikeID:       0,
			SensorUnitID: 0,
		}
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error(), "example": example},
		)
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}

	if err := db.(*gorm.DB).Create(&trip).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, trip)
}
