package mappers

import (
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api/controllers/models"
	"math"

	domainmodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct {
}

// NewBetStorageMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapDomainBetToStorageBet maps the given domain bet into storage bet. Floating point values will
// be converted to corresponding integer values of the storage bet by multiplying them with 100.
func (m *BetMapper) MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.Bet {
	return storagemodels.Bet{
		Id:                   domainBet.Id,
		CustomerId:           domainBet.CustomerId,
		Status:               domainBet.Status,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: int(math.Round(domainBet.SelectionCoefficient * 100)),
		Payment:              int(math.Round(domainBet.Payment * 100)),
		Payout:               int(math.Round(domainBet.Payout * 100)),
	}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet {
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		CustomerId:           storageBet.CustomerId,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}

func (m *BetMapper) MapDomainBetToDtoBet(domainBet domainmodels.Bet) models.BetResponseDto {
	return models.BetResponseDto{
		Id:                   domainBet.Id,
		Status:               domainBet.Status,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: domainBet.SelectionCoefficient,
		Payment:              domainBet.Payment,
		Payout:               domainBet.Payout,
	}
}
