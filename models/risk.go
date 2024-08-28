package models

import "github.com/google/uuid"

// RiskState defines a custom type for the state of a Risk
type RiskState string

// Enum values for RiskState
const (
    Open          RiskState = "open"
    Closed        RiskState = "closed"
    Accepted      RiskState = "accepted"
    Investigating RiskState = "investigating"
)

// Risk represents a risk entity
type Risk struct {
    ID          uuid.UUID  `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    State       RiskState  `json:"state"`
}