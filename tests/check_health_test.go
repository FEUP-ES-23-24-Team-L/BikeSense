package tests

import (
	server "bikesense-web/internal"
	"net/http"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	// launch server
	go server.Run()

	// TestCheckHealth
	resp, err := http.Get("http://localhost:8080/CheckHealth")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", resp.StatusCode)
	}

	t.Log("CheckHealth test passed")

	defer resp.Body.Close()
}
