package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api/controllers/models"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/models"
)

type DtoBetMapper interface {
	MapDomainBetToDtoBet(domainBet domainmodels.Bet) models.BetResponseDto
}
