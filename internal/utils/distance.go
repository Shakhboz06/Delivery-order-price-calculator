package utils

import (
	"fmt"
	"math"
)

func CalDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const Radius = 6371000 // Earth's radius in meters

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Apply the Haversine formula
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := Radius * c

	// Debug logs
	fmt.Printf("DEBUG: Distance calculated: %.2f meters (lat1=%.6f, lon1=%.6f, lat2=%.6f, lon2=%.6f)\n",
		distance, lat1, lon1, lat2, lon2)

	return distance
}
