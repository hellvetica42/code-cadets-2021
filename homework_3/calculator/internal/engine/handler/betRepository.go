package handler

import (
	"context"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.BetCalculated) error
	UpdateBet(ctx context.Context, bet domainmodels.BetCalculated) error
	GetBetBySelectionID(ctx context.Context, id string) ([]domainmodels.BetCalculated, bool, error)
}
