package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
	"log"
)

// BetReceivedConsumer consumes received bets from the desired RabbitMQ queue.
type EventReceivedConsumer struct {
	channel Channel
	config  ConsumerConfig
}

// NewEventReceivedConsumer creates and returns a new EventReceivedConsumer.
func NewEventReceivedConsumer(channel Channel, config ConsumerConfig) (*EventReceivedConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet received consumer initialization failed")
	}

	return &EventReceivedConsumer{
		channel: channel,
		config:  config,
	}, nil
}
// Consume consumes messages until the context is cancelled. An error will be returned if consuming
// is not possible.
func (c *EventReceivedConsumer) Consume(ctx context.Context) (<-chan models.EventReceived, error) {
	msgs, err := c.channel.Consume(
		c.config.Queue,
		c.config.ConsumerName,
		c.config.AutoAck,
		c.config.Exclusive,
		c.config.NoLocal,
		c.config.NoWait,
		c.config.Args,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet received consumer failed to consume messages")
	}

	eventsReceived := make(chan models.EventReceived)
	go func() {
		defer close(eventsReceived)
		for msg := range msgs {
			var eventReceived models.EventReceived
			err := json.Unmarshal(msg.Body, &eventReceived)
			if err != nil {
				log.Println("Failed to unmarshal bet received message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case eventsReceived <- eventReceived:
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventsReceived, nil
}
