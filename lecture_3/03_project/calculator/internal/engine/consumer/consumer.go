package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Consumer offers methods for consuming from input queues.
type Consumer struct {
	betConsumer   BetConsumer
	eventConsumer EventConsumer
}

// New creates and returns a new Consumer.
func New(betConsumer BetConsumer, eventConsumer EventConsumer) *Consumer {
	return &Consumer{
		betConsumer: betConsumer,
		eventConsumer: eventConsumer,
	}
}

// ConsumeBetsReceived consumes bets received queue.
func (c *Consumer) ConsumeEventReceived(ctx context.Context) (<-chan rabbitmqmodels.EventReceived, error) {
	return c.eventConsumer.Consume(ctx)
}

// ConsumeBetsCalculated consumes bets calculated queue.
func (c *Consumer) ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error) {
	return c.betConsumer.Consume(ctx)
}
