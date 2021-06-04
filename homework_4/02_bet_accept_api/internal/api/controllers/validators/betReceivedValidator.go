package validators

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api/controllers/models"

const lostOutcome = "lost"
const wonOutcome = "won"

// BetReceivedValidator validates event update requests.
type BetReceivedValidator struct{}

// BetReceivedUpdateValidator creates a new instance of BetReceivedValidator.
func BetReceivedUpdateValidator() *BetReceivedValidator {
	return &BetReceivedValidator{}
}

// BetReceivedIsValid checks if event update is valid.
// Id is not empty
// Outcome is `lost`or `won`
func (e *BetReceivedValidator) BetReceivedIsValid(betReceivedRequestDto models.BetReceivedRequestDto) bool {
	if betReceivedRequestDto.SelectionCoefficient > 10.0 {
		return false
	}

	if betReceivedRequestDto.Payment < 2.0 || betReceivedRequestDto.Payment > 100.0 {
		return false
	}

	return true
}
