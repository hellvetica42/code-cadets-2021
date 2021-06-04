package models

// BetReceivedRequestDto Update request dto model.
type BetReceivedRequestDto struct {
	CustomerId           string  `form:"customerId"`
	SelectionId          string  `form:"selectionId"`
	SelectionCoefficient float64 `form:"selectionCoefficient"`
	Payment              float64 `form:"payment"`
}
