package main

import (
	"fmt"

	"github.com/shweta-tu/risk-service/api/v1/routes"
	"github.com/shweta-tu/risk-service/storage"
)

func main() {
	store := storage.NewRiskStorage()

	// Setup routes
	router := routes.SetupRoutes(store)

	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Server failed:", err)
	}
}

