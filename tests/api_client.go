package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func testVenueLocation(t *testing.T) {
	// Step 1: Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response for the static endpoint
		if r.URL.Path == "/test-venue/static" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"venue_raw":{"location":{"coordinates":[24.93, 60.17]}}}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Step 2: Hardcode the API URL for testing
	apiURL := fmt.Sprintf("%s/test-venue/static", server.URL)

	// Step 3: Create a function to fetch venue location
	fetchVenueLocation := func(apiURL string) (float64, float64, error) {
		res, err := http.Get(apiURL)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to fetch venue data: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return 0, 0, fmt.Errorf("failed to fetch venue data: Status code: %v", res.StatusCode)
		}

		var venueData struct {
			VenueRaw struct {
				Location struct {
					Coordinates []float64 `json:"coordinates"`
				} `json:"location"`
			} `json:"venue_raw"`
		}

		if err := json.NewDecoder(res.Body).Decode(&venueData); err != nil {
			return 0, 0, fmt.Errorf("failed to decode venue data: %w", err)
		}

		if len(venueData.VenueRaw.Location.Coordinates) != 2 {
			return 0, 0, fmt.Errorf("invalid coordinates in venue data")
		}

		return venueData.VenueRaw.Location.Coordinates[0], venueData.VenueRaw.Location.Coordinates[1], nil
	}

	// Step 4: Call the hardcoded function and test
	long, lat, err := fetchVenueLocation(apiURL)
	assert.NoError(t, err)
	assert.Equal(t, 24.93, long)
	assert.Equal(t, 60.17, lat)
}
