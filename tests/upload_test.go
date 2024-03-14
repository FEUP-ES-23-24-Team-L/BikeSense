package tests

import (
	dbApi "bikesense-web/internal/database"
	"encoding/json"
	"testing"
)

func TestDataInsertion(t *testing.T) {
	jsonData := `{
                    "timestamp": "2023-10-26T19:46:47Z",
                    "location": 55.6842625,
                    "tripID": 1,
                    "noise_level": 2,
                    "temperature": 27,
                    "humidity": 9,
                    "uv_level": 10,
                    "luminosity": 1,
                    "carbon_monoxide": 5,
                    "polution_particles_ppm": 98
                }
  `

	var dataPoint *dbApi.DataPoint
	err := json.Unmarshal([]byte(jsonData), &dataPoint)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Deserialization successfull!")
	t.Log("Timestamp is: ", dataPoint.Timestamp)

	unit := dbApi.SensorUnit{
		Code: "123",
	}

	bike := dbApi.Bike{
		Code: "456",
	}

	db := GetTestDB()

	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Error beginning transaction: %v", tx.Error)
	}

	defer tx.Rollback()

	if err := tx.Create(&unit).Error; err != nil {
		t.Fatalf("Error creating sensor unit: %v", err)
	}

	if err := tx.Create(&bike).Error; err != nil {
		t.Fatalf("Error creating bike: %v", err)
	}

	trip := dbApi.Trip{
		BikeID:       bike.ID,
		SensorUnitID: unit.ID,
	}

	if err := tx.Create(&trip).Error; err != nil {
		t.Fatalf("Error creating trip: %v", err)
	}

	dataPoint.TripID = trip.ID

	if err := tx.Create(&dataPoint).Error; err != nil {
		t.Fatalf("Error creating data point: %v", err)
	}

	var dataPoint2 dbApi.DataPoint
	dataPoint2.ID = dataPoint.ID
	if err := tx.First(&dataPoint2).Error; err != nil {
		t.Fatalf("Error retrieving data point: %v", err)
	}

	if dataPoint2.Location != dataPoint.Location {
		t.Fatalf("Retrieved data point does not match the original. Expected: %v, got: %v",
			dataPoint.Timestamp,
			dataPoint2.Timestamp)
	}
}
