package tests

import (
	dbApi "bikesense-web/internal/database"
	"testing"
	"time"

	"gorm.io/gorm"
)

func insertBike(db *gorm.DB, code string) (uint, error) {
	bike := dbApi.Bike{Code: code}
	err := db.Create(&bike).Error
	id := bike.ID
	return id, err
}

func insertSensorUnit(db *gorm.DB, code string) (uint, error) {
	sensorUnit := dbApi.SensorUnit{Code: code}
	err := db.Create(&sensorUnit).Error
	id := sensorUnit.ID
	return id, err
}

func insertTrip(db *gorm.DB, bikeID, sensorUnitID uint) (uint, error) {
	trip := dbApi.Trip{BikeID: bikeID, SensorUnitID: sensorUnitID}
	err := db.Create(&trip).Error
	id := trip.ID
	return id, err
}

func TestDataInsertion(t *testing.T) {
	db := GetTestDB(t)

	bikeCode := "bike1"
	sensorUnitCode := "sensor1"

	bikeID, err := insertBike(db, bikeCode)
	if err != nil {
		t.Fatalf("Error inserting bike: %v", err)
	}

	sensorUnitID, err := insertSensorUnit(db, sensorUnitCode)
	if err != nil {
		t.Fatalf("Error inserting sensor unit: %v", err)
	}

	tripID, err := insertTrip(db, bikeID, sensorUnitID)
	if err != nil {
		t.Fatalf("Error inserting trip: %v", err)
	}

	dataPoint := dbApi.DataPoint{
		GPSData: dbApi.GPSData{
			Latitude:        200,
			Longitude:       100,
			Speed:           230,
			Altitute:        30,
			Course:          23,
			SatellitesInUse: 1,
			FixType:         2,
			HDOP:            1,
			VDOP:            1,
			PDOP:            1,
		},
		ID:                   bikeID,
		TripID:               tripID,
		NoiseLevel:           2,
		Temperature:          320,
		Humidity:             123,
		UVLevel:              2323,
		Luminosity:           32,
		CarbonMonoxideLevel:  123,
		PolutionParticlesPPM: 12,
	}

	dataPoint.Timestamp = time.Now()

	err = db.Create(&dataPoint).Error
	if err != nil {
		t.Fatal(err)
	}
}
