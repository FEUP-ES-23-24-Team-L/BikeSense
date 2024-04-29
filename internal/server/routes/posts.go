package routes

import (
	"net/http"

	dbApi "bikesense-web/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func bindAndCreateRecord(c *gin.Context, record_type any) {
	bound_record, err := bindRecord(c, record_type)
	if err != nil {
		return
	}

	createRecord(c, bound_record)
}

func bindRecord(c *gin.Context, record_type any) (any, error) {
	if err := c.ShouldBindJSON(&record_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		return nil, err
	}

	return record_type, nil
}

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

func PostSensorUnit(c *gin.Context) {
	bindAndCreateRecord(c, &dbApi.SensorUnit{})
}

func PostBike(c *gin.Context) {
	bindAndCreateRecord(c, &dbApi.Bike{})
}

func PostTrip(c *gin.Context) {
	bindAndCreateRecord(c, &dbApi.Trip{})
}

func PostTripData(c *gin.Context) {
	// bindAndCreateRecord(c, &dbApi.DataPoint{})
	var data []dbApi.DataPoint
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
	}

	createRecord(c, data)

	// Convert raw data to DataPoint
	// data := make([]dbApi.DataPoint, len(rawData))
	// for i, rawPoint := range rawData {
	// 	data[i] = rawPoint.intoDataPoint()
	// }
}
