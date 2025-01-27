package services

import (
	"dopc/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseAPIURL = "https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues"

type Location struct {
	Coordinates []float64 `json:"coordinates"` // [longitude, latitude]
}

// VenueStaticData represents the response from the `/static` endpoint.
type VenueStaticData struct {
	VenueRaw struct {
		Location struct {
			Coordinates []float64 `json:"coordinates"` // [longitude, latitude]
		} `json:"location"`
	} `json:"venue_raw"`
}


func VenueLocation(venueSlug string) (float64, float64, error) {
	apiUrl := fmt.Sprintf("https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/%s/static", venueSlug)

	res, err := http.Get(apiUrl)
	if err != nil{
		return 0, 0, fmt.Errorf("failed to fetch the venue data: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("failed to fetch venue data: Status code: %v", res.StatusCode)
	}

	var venueData VenueStaticData

	if err := json.NewDecoder(res.Body).Decode(&venueData); err != nil {
		return 0, 0, fmt.Errorf("failed to decode venue data: %w", err)
	}

	// Extract coordinates (longitude, latitude)
	if len(venueData.VenueRaw.Location.Coordinates) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinates in venue data: %f", venueData.VenueRaw.Location.Coordinates)
	}

	longitude := venueData.VenueRaw.Location.Coordinates[0]
	latitude := venueData.VenueRaw.Location.Coordinates[1]

	return longitude, latitude, nil

}



func DynamicPricing(venueSlug string) (models.DynamicPricingData, error) {
	apiURL := fmt.Sprintf("https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/%s/dynamic", venueSlug)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return models.DynamicPricingData{}, fmt.Errorf("failed to fetch dynamic pricing: %w", err)
	}
	defer resp.Body.Close()

	// Check if the status code is OK
	if resp.StatusCode != http.StatusOK {
		return models.DynamicPricingData{}, fmt.Errorf("failed to fetch dynamic pricing: Status code: %v", resp.StatusCode)
	}

	// Decode the JSON directly from the response body
	var pricingData models.DynamicPricingData
	if err := json.NewDecoder(resp.Body).Decode(&pricingData); err != nil {
		return models.DynamicPricingData{}, fmt.Errorf("failed to decode dynamic pricing data: %w", err)
	}

	return pricingData, nil
}


