package tests

import (
	dbApi "bikesense-web/internal/database"
	"testing"
)

func TestParseHHMMSS(t *testing.T) {
	timeStr := "123456789"
	_, err := dbApi.ParseHHMMSS(timeStr)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	timeStr = "123456"
	time, err := dbApi.ParseHHMMSS(timeStr)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if time.Hour() != 12 {
		t.Fatalf("Expected 12, got %v", time.Hour())
	}

	if time.Minute() != 34 {
		t.Fatalf("Expected 34, got %v", time.Minute())
	}

	if time.Second() != 56 {
		t.Fatalf("Expected 56, got %v", time.Second())
	}

	timeStr = "123456.78"
	time, err = dbApi.ParseHHMMSS(timeStr)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if time.Hour() != 12 {
		t.Fatalf("Expected 12, got %v", time.Hour())
	}

	if time.Minute() != 34 {
		t.Fatalf("Expected 34, got %v", time.Minute())
	}

	if time.Second() != 56 {
		t.Fatalf("Expected 56, got %v", time.Second())
	}
}

func TestGPGGADecode(t *testing.T) {
	message := "GPGGA,123456,123.678,N,123.678,W,1,08,0.9,545.4,M,46.9,M,,"
	data, err := dbApi.DecodeGPGGA(message)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if data.Timestamp.Hour() != 12 {
		t.Fatalf("Expected 12, got %v", data.Timestamp.Hour())
	}
	if data.Timestamp.Minute() != 34 {
		t.Fatalf("Expected 34, got %v", data.Timestamp.Minute())
	}
	if data.Timestamp.Second() != 56 {
		t.Fatalf("Expected 56, got %v", data.Timestamp.Second())
	}
	if data.Latitude != 123.678 {
		t.Fatalf("Expected 123.678, got %v", data.Latitude)
	}
	if data.Longitude != -123.678 {
		t.Fatalf("Expected -123.678, got %v", data.Longitude)
	}
}
