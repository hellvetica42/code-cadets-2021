package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Consumer interface {
	ConsumeBetsReceived(ctx context.Context) (<-chan rabbitmqmodels.Bet, error)
	ConsumeEventsReceived (ctx context.Context) (<-chan rabbitmqmodels.EventReceived, error)
}
