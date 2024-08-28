package storage

import (
	"sync"

	"github.com/google/uuid"
	"github.com/shweta-tu/risk-service/models"
)

// RiskStorage represents the in-memory storage of risks
type RiskStorage struct {
    Risks []models.Risk
    mu    sync.Mutex
}

// NewRiskStorage creates a new in-memory storage for risks
func NewRiskStorage() *RiskStorage {
    return &RiskStorage{
        Risks: make([]models.Risk, 0),
    }
}

// AddRisk adds a new risk to the storage
func (s *RiskStorage) AddRisk(risk models.Risk) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.Risks = append(s.Risks, risk)
}

// GetRisks returns all the risks from storage
func (s *RiskStorage) GetRisks() []models.Risk {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.Risks
}

// GetRiskByID returns a specific risk by ID
func (s *RiskStorage) GetRiskByID(id uuid.UUID) (*models.Risk, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, risk := range s.Risks {
        if risk.ID == id {
            return &risk, true
        }
    }
    return nil, false
}
