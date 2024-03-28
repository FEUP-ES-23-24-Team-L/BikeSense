package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseHHMMSS(timeStr string) (time.Time, error) {
	fields := strings.Split(timeStr, ".")

	if len(fields) > 2 || len(fields) < 1 {
		return time.Time{}, fmt.Errorf("invalid time format. Expected HHMMSS.ss, got %v", timeStr)
	}

	milliseconds := 0
	timeStr = fields[0]
	if len(timeStr) != 6 {
		return time.Time{}, fmt.Errorf("invalid time format. Expected HHMMSS.ss, got %v", timeStr)
	}

	// Extract hours, minutes, seconds, and milliseconds
	hours, err := strconv.Atoi(timeStr[:2])
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing hours: %w", err)
	}

	minutes, err := strconv.Atoi(timeStr[2:4])
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing minutes: %w", err)
	}

	seconds, err := strconv.Atoi(timeStr[4:6])
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing seconds: %w", err)
	}

	if len(fields) == 2 {
		milliseconds, err = strconv.Atoi(fields[1])
		if err != nil {
			return time.Time{}, fmt.Errorf("error parsing milliseconds: %w", err)
		}
	}

	// Create a time.Time object
	return time.Date(0, 1, 1, hours, minutes, seconds, milliseconds*int(time.Millisecond), time.UTC), nil
}

func DecodeGPGGA(message string) (*GPGGAData, error) {
	var err error

	fields := strings.Split(message, ",")
	if len(fields) < 4 {
		return nil, fmt.Errorf("invalid GPGGA message format")
	}

	data := &GPGGAData{}

	// Parse timestamp (assuming field 1)
	data.Timestamp, err = ParseHHMMSS(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp %v", fields[1])
	}

	latitudeStr := fields[2]
	data.Latitude, err = strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing latitude: %w", err)
	}

	longitudeStr := fields[4]
	data.Longitude, err = strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing longitude degrees: %w", err)
	}

	// Check for north/south and east/west indicators
	if fields[3] == "S" {
		data.Latitude *= -1
	}

	if fields[5] == "W" {
		data.Longitude *= -1
	}

	// Parse optional fields (fix quality, satellites used, altitude)
	if len(fields) > 6 {
		fixQuality, err := strconv.Atoi(fields[6])
		if err != nil {
			return nil, fmt.Errorf("error parsing fix quality: %w", err)
		}
		data.FixQuality = fixQuality
	}

	if len(fields) > 7 {
		satellitesUsed, err := strconv.Atoi(fields[7])
		if err != nil {
			return nil, fmt.Errorf("error parsing satellites used: %w", err)
		}
		data.SatellitesUsed = satellitesUsed
	}

	if len(fields) > 9 {
		altitude, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing altitude: %w", err)
		}
		data.Altitude = altitude
	}

	return data, nil
}
