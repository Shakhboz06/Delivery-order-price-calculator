package tests

import (
	"dopc/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVenueLocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/home-assignment-api/v1/venues/test-venue/static" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"venue_raw":{"location":{"coordinates":[13.453615, 52.50032]}}}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	services.BaseAPIURL = server.URL

	lon, lat, err := services.VenueLocation("test-venue")
	assert.NoError(t, err)
	assert.Equal(t, 13.453615, lon)
	assert.Equal(t, 52.50032, lat)
}
