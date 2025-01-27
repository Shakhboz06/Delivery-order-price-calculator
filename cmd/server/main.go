package main

import (
	"dopc/internal/handlers"
	"dopc/internal/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	delService := services.CalculatorService()
	delHandler := handlers.DeliveryHandler(*delService)
	// Create a Gin router
	router := gin.Default()

	// Define a placeholder route
	router.GET("/api/v1/delivery-order-price", delHandler.DeliveryPrice)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

}
