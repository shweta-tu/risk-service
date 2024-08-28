package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shweta-tu/risk-service/handlers"
	"github.com/shweta-tu/risk-service/storage"
)

// SetupRoutes initializes the routes for the API using Gin
func SetupRoutes(store *storage.RiskStorage) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/v1/risks", handlers.HandleGetRisks(store))
	r.GET("/v1/risks/:id", handlers.HandleGetRiskByID(store))
	r.POST("/v1/risks", handlers.HandleCreateRisk(store))

	return r
}
