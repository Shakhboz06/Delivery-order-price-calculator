package models

type DynamicPricingData struct {
    VenueRaw struct {
        DeliverySpecs struct {
            OrderMinimumNoSurcharge int `json:"order_minimum_no_surcharge"` // Minimum order value for no surcharge
            DeliveryPricing         struct {
                BasePrice      int `json:"base_price"` // Base delivery fee
                DistanceRanges []struct {
                    Min int     `json:"min"` // Minimum distance for the range
                    Max int     `json:"max"` // Maximum distance for the range (0 = no limit)
                    A   int     `json:"a"`   // Fixed surcharge for the range
                    B   float64 `json:"b"`   // Multiplier for distance-based fee
                } `json:"distance_ranges"` // Ranges for calculating distance-based fees
            } `json:"delivery_pricing"`
        } `json:"delivery_specs"`
    } `json:"venue_raw"`
}