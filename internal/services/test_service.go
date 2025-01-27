package services

import "dopc/internal/models"

type MockCalculateService struct {
    CalculateDeliveryPriceFunc func(venueSlug string, cartValue int, userLat, userLon float64) (models.Response, error)
}

func (m *MockCalculateService) CalculateDeliveryPrice(venueSlug string, cartValue int, userLat, userLon float64) (models.Response, error) {
    if m.CalculateDeliveryPriceFunc != nil {
        return m.CalculateDeliveryPriceFunc(venueSlug, cartValue, userLat, userLon)
    }
    return models.Response{}, nil
}
