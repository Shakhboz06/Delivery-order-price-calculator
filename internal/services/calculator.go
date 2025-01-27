package services

import (
	"dopc/internal/models"
	"dopc/internal/utils"
	"errors"
	"fmt"
)

// Service defines the interface for calculating delivery prices
type Service interface {
	CalculateDeliveryPrice(venueSlug string, cartValue int, userLat, userLon float64) (models.Response, error)
}

// CalculateService is the implementation of the Service interface
type CalculateService struct{}

// CalculatorService initializes and returns a new CalculateService instance
func CalculatorService() *CalculateService {
	return &CalculateService{}
}

// CalculateDeliveryPrice computes the delivery price based on venue details, user location, and cart value
func (s *CalculateService) CalculateDeliveryPrice(venueSlug string, cartValue int, userLat, userLon float64) (models.Response, error) {
	if venueSlug == "" {
		return models.Response{}, errors.New("venue_slug cannot be empty")
	}

	// Fetch venue location
	venueLon, venueLat, err := VenueLocation(venueSlug)
	if err != nil {
		return models.Response{}, err
	}

	// Fetch dynamic pricing details
	dynamicPrice, err := DynamicPricing(venueSlug)
	if err != nil {
		return models.Response{}, err
	}

	// Calculate the distance between user and venue
	distance := utils.CalDistance(userLat, userLon, venueLat, venueLon)
	fmt.Printf("DEBUG: Calculated distance = %.2f meters\n", distance)

	// Calculate the delivery fee
	deliveryFee, err := calculateDeliveryFee(distance, dynamicPrice.VenueRaw.DeliverySpecs.DeliveryPricing)
	fmt.Printf("DEBUG: Calculated delivery fee = %d cents\n", deliveryFee)
	if err != nil {
		return models.Response{}, fmt.Errorf("delivery not available for the given distance: %w", err)
	}

	// Calculate small order surcharge
	orderMinimum := dynamicPrice.VenueRaw.DeliverySpecs.OrderMinimumNoSurcharge
	smallOrderSurcharge := max(0, orderMinimum-cartValue)

	// Calculate the total price
	totalPrice := cartValue + deliveryFee + smallOrderSurcharge

	// Build the response
	response := models.Response{
		TotalPrice:          totalPrice,
		SmallOrderSurcharge: smallOrderSurcharge,
		CartValue:           cartValue,
		Delivery: struct {
			Fee      int `json:"fee"`
			Distance int `json:"distance"`
		}{
			Fee:      deliveryFee,
			Distance: int(distance),
		},
	}

	return response, nil
}

// calculateDeliveryFee calculates the delivery fee based on distance and pricing rules
func calculateDeliveryFee(distance float64, pricing struct {
	BasePrice      int `json:"base_price"`
	DistanceRanges []struct {
		Min int     `json:"min"`
		Max int     `json:"max"`
		A   int     `json:"a"`
		B   float64 `json:"b"`
	} `json:"distance_ranges"`
}) (int, error) {
	if distance < 10 {
		return 0, fmt.Errorf("delivery not available for distances less than 10 meters")
	}
	for _, rangeData := range pricing.DistanceRanges {
		// If Max = 0 and distance is greater than or equal to Min, delivery is unavailable
		if rangeData.Max == 0 && int(distance) >= rangeData.Min {
			return 0, fmt.Errorf("delivery not available for given distance.")
		}

		// Check if the distance falls within the range
		if int(distance) >= rangeData.Min && (rangeData.Max == 0 || int(distance) < rangeData.Max) {
			// Fee = base price + a + (b * distance / 10)
			distanceComponent := int(rangeData.B * (distance / 10))
			return pricing.BasePrice + rangeData.A + distanceComponent, nil
		}
	}

	return 0, fmt.Errorf("delivery not available for the given distance")
}

