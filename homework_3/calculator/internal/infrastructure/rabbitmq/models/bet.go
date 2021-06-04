package models

// Bet represents a DTO for bets.
type Bet struct {
	Id                   string  `json:"id"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}
