package models

// BetCalculated represents the calculated status of a bet
type BetCalculated struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
}
