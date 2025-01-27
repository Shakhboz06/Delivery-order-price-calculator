package handlers

import (
	"dopc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service services.CalculateService
}

func DeliveryHandler(service services.CalculateService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) DeliveryPrice(ctx *gin.Context) {
	
	venueSlug := ctx.Query("venue_slug")
	cartValueStr := ctx.Query("cart_value")
	userLatStr := ctx.Query("user_lat")
	userLonStr := ctx.Query("user_lon")


	if venueSlug == "" || cartValueStr == "" || userLatStr == "" || userLonStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameters"})
		return
	}

	cartValue, err := strconv.Atoi(cartValueStr)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart value"})
		return
	}

	userLat, err := strconv.ParseFloat(userLatStr, 64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user latitude"})
		return
	}

	userLon, err := strconv.ParseFloat(userLonStr, 64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user longtitude"})
		return
	}

	
	result, err := h.Service.CalculateDeliveryPrice(venueSlug, cartValue, userLat, userLon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, result)

}
