package models

// Bet is a storage model representation of a bet.
type BetCalculated struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
}
