package controllers

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/models"
)


type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateBet(ctx context.Context, bet domainmodels.Bet) error
	GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error)
	GetBetsByUserID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error)
	GetBetsByStatus(ctx context.Context, id string) ([]domainmodels.Bet, bool, error)
}
