package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shweta-tu/risk-service/models"
	"github.com/shweta-tu/risk-service/storage"
	"github.com/shweta-tu/risk-service/utils"
)

// HandleGetRisks handles GET requests to /v1/risks
func HandleGetRisks(store *storage.RiskStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		risks := store.GetRisks()
		utils.RespondWithJSON(c, http.StatusOK, risks)
	}
}

// HandleCreateRisk handles POST requests to /v1/risks
func HandleCreateRisk(store *storage.RiskStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRisk models.Risk
		if err := c.ShouldBindJSON(&newRisk); err != nil {
			utils.LogError(err)
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		switch newRisk.State {
		case models.Open, models.Closed, models.Accepted, models.Investigating:
			// valid state
		default:
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid state value")
			return
		}

		newRisk.ID = uuid.New()
		store.AddRisk(newRisk)

		utils.RespondWithJSON(c, http.StatusCreated, newRisk)
	}
}

// HandleGetRiskByID handles GET requests to /v1/risks/<id>
func HandleGetRiskByID(store *storage.RiskStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			utils.LogError(err)
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid risk ID")
			return
		}

		risk, found := store.GetRiskByID(id)
		if !found {
			utils.RespondWithError(c, http.StatusNotFound, "Risk not found")
			return
		}

		utils.RespondWithJSON(c, http.StatusOK, risk)
	}
}
