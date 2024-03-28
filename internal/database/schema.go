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
	GPGGAData            `json:"gpgga_data"`
	ID                   uint    `json:"id" gorm:"primaryKey"`
	TripID               uint    `json:"trip_id"`
	NoiseLevel           float32 `json:"noise_level"`
	Temperature          float32 `json:"temperature"`
	Humidity             float32 `json:"humidity"`
	UVLevel              float32 `json:"uv_level"`
	Luminosity           float32 `json:"luminosity"`
	CarbonMonoxideLevel  float32 `json:"carbon_monoxide_level"`
	PolutionParticlesPPM int32   `json:"polution_particles_ppm"`
}

type GPGGAData struct {
	Timestamp      time.Time `json:"timestamp"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	FixQuality     int       `json:"fix_quality"`
	SatellitesUsed int       `json:"satellites_used"`
	Altitude       float64   `json:"altitude"`
}
