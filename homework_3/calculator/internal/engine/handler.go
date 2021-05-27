package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBetsReceived(ctx context.Context, betsReceived <-chan rabbitmqmodels.Bet)
	HandleEventsReceived(ctx context.Context, betsCalculated <-chan rabbitmqmodels.EventReceived) <-chan rabbitmqmodels.BetCalculated
}
