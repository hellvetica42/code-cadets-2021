package validators

import "regexp"

const lostOutcome = "lost"
const wonOutcome = "won"

// EventUpdateValidator validates event update requests.
type BetRequestValidator struct{}

// NewEventUpdateValidator creates a new instance of EventUpdateValidator.
func NewBetRequestValidator() *BetRequestValidator {
	return &BetRequestValidator{}
}

func (* BetRequestValidator) BetRequestIdIsValid(id string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	return r.MatchString(id)
}
func (* BetRequestValidator) BetRequestUserIdIsValid(id string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	return r.MatchString(id)
}
func (* BetRequestValidator) BetRequestStatusIsValid(status string) bool {
	if status == "won" || status == "lost" || status == "active" {
		return true
	} else {
		return false
	}
}
