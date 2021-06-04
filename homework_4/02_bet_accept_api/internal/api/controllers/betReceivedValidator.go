package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api/controllers/models"

// BetReceivedValidator validates event update requests.
type BetReceivedValidator interface {
	BetReceivedIsValid(betReceivedRequestDto models.BetReceivedRequestDto) bool
}
