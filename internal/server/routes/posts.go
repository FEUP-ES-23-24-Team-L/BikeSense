package routes

import (
	"net/http"

	dbApi "bikesense-web/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostSensorUnit(c *gin.Context) {
	var sensorUnit dbApi.SensorUnit

	if err := c.ShouldBindJSON(&sensorUnit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}

	if err := db.(*gorm.DB).Create(&sensorUnit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sensorUnit)
}

func PostBike(c *gin.Context) {
	var bike dbApi.Bike

	if err := c.ShouldBindJSON(&bike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}

	if err := db.(*gorm.DB).Create(&bike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bike)
}

func PostTrip(c *gin.Context) {
	var trip dbApi.Trip

	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
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

func PostTripData(c *gin.Context) {
	var tripData []dbApi.DataPoint
	if err := c.ShouldBindJSON(&tripData); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}
	if err := db.(*gorm.DB).Create(&tripData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tripData)
}
