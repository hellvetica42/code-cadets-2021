package services

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq/models"

// BetReceivedService implements event related functions.
type BetReceivedService struct {
	betReceivedPublisher BetReceivedPublisher
}

// NewBetReceivedService creates a new instance of BetReceivedService.
func NewBetReceivedService(betReceivedPublisher BetReceivedPublisher) *BetReceivedService {
	return &BetReceivedService{
		betReceivedPublisher: betReceivedPublisher,
	}
}

// UpdateBetReceived sends event update message to the queues.
func (e BetReceivedService) PublishBet(betReceivedDto models.BetReceivedDto) error {
	return e.betReceivedPublisher.Publish(betReceivedDto)
}
