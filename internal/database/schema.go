package database

import (
	"time"

	"gorm.io/gorm"
)

type Bike struct {
	gorm.Model
	Code string
	trip Trip
}

type SensorUnit struct {
	gorm.Model
	Code string
	trip Trip
}

type Trip struct {
	Time time.Time
	gorm.Model
	DataPoint      []DataPoint
	Duration       time.Duration
	BikeID         uint
	SensorUnitID   uint
	TravelDistance float32
}

type DataPoint struct {
	Timestamp time.Time `json:"timestamp"`
	gorm.Model
	TripID               uint
	Location             float32 `json:"location"`
	NoiseLevel           float32 `json:"noise_level"`
	Temperature          float32 `json:"temperature"`
	Humidity             float32 `json:"humidity"`
	UVLevel              float32 `json:"uv_level"`
	Luminosity           float32 `json:"luminosity"`
	CarbonMonoxideLevel  float32 `json:"carbon_monoxide_level"`
	PolutionParticlesPPM int32   `json:"polution_particles_ppm"`
}
