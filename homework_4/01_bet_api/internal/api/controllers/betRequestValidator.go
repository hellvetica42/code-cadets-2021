package controllers

// BetRequestValidator validates event update requests.
type BetRequestValidator interface {
	BetRequestIdIsValid(id string) bool
	BetRequestUserIdIsValid(id string) bool
	BetRequestStatusIsValid(status string) bool
}
