package database

import (
	"time"
)

type Bike struct {
	Code string `json:"code"`
	trip Trip   `json:"-"`
	ID   uint   `json:"id" gorm:"primaryKey"`
}

type SensorUnit struct {
	Code string `json:"code"`
	Trip Trip   `json:"-"`
	ID   uint   `json:"id" gorm:"primaryKey"`
}

type Trip struct {
	DataPoint    []DataPoint `json:"-"`
	ID           uint        `json:"id" gorm:"primaryKey"`
	BikeID       uint        `json:"bike_id"`
	SensorUnitID uint        `json:"sensor_unit_id"`
}

type DataPoint struct {
	Timestamp            time.Time `json:"timestamp"`
	ID                   uint      `json:"id" gorm:"primaryKey"`
	TripID               uint      `json:"trip_id"`
	Location             float32   `json:"location"`
	NoiseLevel           float32   `json:"noise_level"`
	Temperature          float32   `json:"temperature"`
	Humidity             float32   `json:"humidity"`
	UVLevel              float32   `json:"uv_level"`
	Luminosity           float32   `json:"luminosity"`
	CarbonMonoxideLevel  float32   `json:"carbon_monoxide_level"`
	PolutionParticlesPPM int32     `json:"polution_particles_ppm"`
}
