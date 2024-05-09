package routes

import (
	"net/http"
	"strconv"

	dbApi "bikesense-web/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createRecord(c *gin.Context, record any) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}

	if err := db.(*gorm.DB).Create(record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DB error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func createUniqueRecord(c *gin.Context, record any, cond any) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection not found in context"})
		return
	}

	if err := db.(*gorm.DB).FirstOrCreate(record, cond).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DB error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func PostSensorUnit(c *gin.Context) {
	var unit dbApi.SensorUnit
	err := c.ShouldBindJSON(&unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		return
	}

	createUniqueRecord(c, &unit, &dbApi.SensorUnit{Code: unit.Code})
}

func PostBike(c *gin.Context) {
	var bike dbApi.Bike
	err := c.ShouldBindJSON(&bike)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		return
	}

	createUniqueRecord(c, &bike, &dbApi.Bike{Code: bike.Code})
}

func PostTrip(c *gin.Context) {
	var trip dbApi.Trip
	err := c.ShouldBindJSON(&trip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		return
	}

	createRecord(c, &trip)
}

func PostTripData(c *gin.Context) {
	tripIdStr := c.GetHeader("Trip-ID")
	if tripIdStr == "" {
		c.JSON(http.StatusBadRequest, "Missing 'Trip-ID' header")
		return
	}

	tripId, err := strconv.Atoi(tripIdStr)
	if err != nil || tripId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid trip id": tripIdStr})
		return
	}

	var data []dbApi.DataPoint
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		return
	}

	for i := 0; i < len(data); i++ {
		data[i].TripID = uint(tripId)
	}

	createRecord(c, &data)
}
