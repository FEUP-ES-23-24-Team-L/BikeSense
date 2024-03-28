package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

type rawTripDataPoint struct {
	Timestamp            time.Time `json:"timestamp"`
	GPGGAMessage         string    `json:"gpgga_message"`
	ID                   uint      `json:"id" gorm:"primaryKey"`
	TripID               uint      `json:"trip_id"`
	NoiseLevel           float32   `json:"noise_level"`
	Temperature          float32   `json:"temperature"`
	Humidity             float32   `json:"humidity"`
	UVLevel              float32   `json:"uv_level"`
	Luminosity           float32   `json:"luminosity"`
	CarbonMonoxideLevel  float32   `json:"carbon_monoxide_level"`
	PolutionParticlesPPM int32     `json:"polution_particles_ppm"`
}

// TODO: Return error if location string is not in the correct format
func (data *rawTripDataPoint) intoDataPoint() dbApi.DataPoint {
	// Parse location string
	gpgga, err := dbApi.DecodeGPGGA(data.GPGGAMessage)
	if err != nil {
		log.Println(fmt.Errorf("error decoding GPGGA message: %v", err))
		return dbApi.DataPoint{}
	}

	return dbApi.DataPoint{
		GPGGAData:            *gpgga,
		ID:                   0,
		TripID:               data.TripID,
		NoiseLevel:           data.NoiseLevel,
		Temperature:          data.Temperature,
		Humidity:             data.Humidity,
		UVLevel:              data.UVLevel,
		Luminosity:           data.Luminosity,
		CarbonMonoxideLevel:  data.CarbonMonoxideLevel,
		PolutionParticlesPPM: data.PolutionParticlesPPM,
	}
}

func PostTripData(c *gin.Context) {
	var rawData []rawTripDataPoint
	if err := c.ShouldBindJSON(&rawData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
	}

	// Convert raw data to DataPoint
	data := make([]dbApi.DataPoint, len(rawData))
	for i, rawPoint := range rawData {
		data[i] = rawPoint.intoDataPoint()
	}

	createRecord(c, data)
}
