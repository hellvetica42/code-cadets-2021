package rabbitmq

import (
	"encoding/json"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq/models"
	"log"

	"github.com/streadway/amqp"
)

const contentTypeTextPlain = "text/plain"

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetReceivedPublisher create a new instance of BetReceivedPublisher.
func NewBetReceivedPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetReceivedPublisher {
	return &BetReceivedPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes an event update message to the queue.
func (p *BetReceivedPublisher) Publish(betReceivedDto models.BetReceivedDto) error {

	betReceived, err := json.Marshal(betReceivedDto)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betReceived,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betReceived)
	return nil
}
