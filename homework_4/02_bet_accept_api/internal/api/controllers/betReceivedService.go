package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq/models"
)

// BetReceivedService implements event related functions.
type BetReceivedService interface {
	PublishBet(betReceivedDto models.BetReceivedDto) error
}
