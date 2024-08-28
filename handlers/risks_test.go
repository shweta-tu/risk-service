package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shweta-tu/risk-service/models"
	"github.com/shweta-tu/risk-service/storage"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetRisks(t *testing.T) {
	store := storage.NewRiskStorage()
	r := gin.Default()
	r.GET("/v1/risks", HandleGetRisks(store))

	// Test empty storage
	req, _ := http.NewRequest("GET", "/v1/risks", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]", rec.Body.String())

	// Test with a risk
	risk := models.Risk{
		ID:          uuid.New(),
		State:       models.Open,
		Title:       "Test Risk",
		Description: "Test Description",
	}
	store.AddRisk(risk)
	req, _ = http.NewRequest("GET", "/v1/risks", nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var risks []models.Risk
	err := json.Unmarshal(rec.Body.Bytes(), &risks)
	assert.Nil(t, err)
	assert.Len(t, risks, 1)
	assert.Equal(t, risk, risks[0])
}

func TestHandleCreateRisk(t *testing.T) {
	store := storage.NewRiskStorage()
	r := gin.Default()
	r.POST("/v1/risks", HandleCreateRisk(store))

	// Test successful creation
	risk := models.Risk{
		State:       models.Open,
		Title:       "New Risk",
		Description: "Description",
	}
	body, _ := json.Marshal(risk)
	req, _ := http.NewRequest("POST", "/v1/risks", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var createdRisk models.Risk
	err := json.Unmarshal(rec.Body.Bytes(), &createdRisk)
	assert.Nil(t, err)
	assert.Equal(t, risk.Title, createdRisk.Title)
	assert.Equal(t, risk.Description, createdRisk.Description)

	// Test invalid JSON
	invalidBody := `{"state": "open", "title": "Invalid Risk"`
	req, _ = http.NewRequest("POST", "/v1/risks", bytes.NewReader([]byte(invalidBody)))
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Test invalid state
	invalidRisk := models.Risk{
		State:       "invalid_state",
		Title:       "Invalid Risk",
		Description: "Description",
	}
	body, _ = json.Marshal(invalidRisk)
	req, _ = http.NewRequest("POST", "/v1/risks", bytes.NewReader(body))
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleGetRiskByID(t *testing.T) {
	store := storage.NewRiskStorage()
	r := gin.Default()
	r.GET("/v1/risks/:id", HandleGetRiskByID(store))

	// Test valid ID
	riskID := uuid.New()
	risk := models.Risk{
		ID:          riskID,
		State:       models.Open,
		Title:       "Test Risk",
		Description: "Test Description",
	}
	store.AddRisk(risk)
	req, _ := http.NewRequest("GET", "/v1/risks/"+riskID.String(), nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var fetchedRisk models.Risk
	err := json.Unmarshal(rec.Body.Bytes(), &fetchedRisk)
	assert.Nil(t, err)
	assert.Equal(t, risk, fetchedRisk)

	// Test invalid ID
	req, _ = http.NewRequest("GET", "/v1/risks/invalid_id", nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Test non-existent ID
	req, _ = http.NewRequest("GET", "/v1/risks/"+uuid.New().String(), nil)
	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
