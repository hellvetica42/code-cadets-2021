package handler

import (
	"context"
	"log"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets received and bets calculated.
type Handler struct {
	betRepository BetRepository
}

// New creates and returns a new Handler.
func New(betRepository BetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

// HandleBetsReceived handles bets received.
func (h *Handler) HandleEventsReceived(
	ctx context.Context,
	eventsReceived <-chan rabbitmqmodels.EventReceived,
) <-chan rabbitmqmodels.BetCalculated{
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventReceived := range eventsReceived {
			log.Println("Processing event received, eventId:", eventReceived.Id)

			//get bet from calc_bets db
			domainBets, exists, err := h.betRepository.GetBetBySelectionID(ctx, eventReceived.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("A bet which should be updated does not exist, betId: ", eventReceived.Id)
				continue
			}

			for _, bet := range domainBets {

				var payout float64
				if eventReceived.Outcome == "won" {
					payout = bet.Payment * bet.SelectionCoefficient
				}else{
					payout = 0
				}
				//create rabbitmq DTO for bets calculated
				betCalculated := rabbitmqmodels.BetCalculated{
					Id:     bet.Id,
					Status: eventReceived.Outcome,
					Payout: payout,
				}

				log.Println("Calculated bet id:", bet.Id, "outcome:", eventReceived.Outcome, "payment:", bet.Payment,
								"coef:", bet.SelectionCoefficient, "payout:", payout)

				//send to calculated bets channel
				select {
				case resultingBets <- betCalculated:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return resultingBets
}

// HandleBetsReceived stores received bets in calc_bets repository
func (h *Handler) HandleBetsReceived(
	ctx context.Context,
	betsReceived <-chan rabbitmqmodels.Bet,
) {

	go func() {

		for betReceived := range betsReceived{
			log.Println("Processing bet received, betId:", betReceived.Id)

			//Creates domain bet that gets stored in sql
			bet := domainmodels.BetCalculated{
				Id:                   betReceived.Id,
				SelectionId:          betReceived.SelectionId,
				SelectionCoefficient: betReceived.SelectionCoefficient,
				Payment:              betReceived.Payment,
			}

			err := h.betRepository.InsertBet(ctx, bet)
			if err != nil {
				err := h.betRepository.UpdateBet(ctx, bet)
				if err != nil {
					log.Println("Failed to insert bet, error: ", err)
					continue
				}
			}

		}
		select {
		case <-ctx.Done():
			return
		}
	}()

}
