package services

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq/models"

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher interface {
	Publish(betReceivedDto models.BetReceivedDto) error
}
