package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// InsertBet inserts the provided bet into the database. An error is returned if the operation
// has failed.
func (r *BetRepository) InsertBet(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	// If payout is 0, do not insert it.

	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBet updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *BetRepository) UpdateBet(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetBySelectionID(ctx context.Context, id string) ([]domainmodels.BetCalculated, bool, error) {
	storageBet, err := r.queryGetBetBySelectionID(ctx, id)
	if err == sql.ErrNoRows {
		return []domainmodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetCalculated{}, false, errors.Wrap(err, "bet repository failed to get a bets with selection id "+id)
	}

	var bets []domainmodels.BetCalculated
	for _, bet := range storageBet {
		domainBet := r.betMapper.MapStorageBetToDomainBet(bet)
		bets = append(bets, domainBet)
	}

	return bets, true, nil
}

func (r *BetRepository) queryGetBetBySelectionID(ctx context.Context, selectionId string) ([]storagemodels.BetCalculated, error) {
	rows, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+selectionId+"';")
	if err != nil {
		return []storagemodels.BetCalculated{}, err
	}
	defer rows.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	var bets []storagemodels.BetCalculated
	for rows.Next() {
		var id string
		var selectionCoefficient int
		var payment int

		err = rows.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return []storagemodels.BetCalculated{}, err
		}

		bet := storagemodels.BetCalculated{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		}

		bets = append(bets, bet)
	}

	return bets, nil
}
