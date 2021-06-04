package sqlite

import (
	domainmodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.Bet
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet
}
