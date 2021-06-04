package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/infrastructure/sqlite/models"
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
func (r *BetRepository) InsertBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.Bet) error {
	// If payout is 0, do not insert it.
	if bet.Payout == 0 {
		insertBetSQL := "INSERT INTO bets(id, customer_id, status, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?, ?, ?)"
		statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
		if err != nil {
			return err
		}

		_, err = statement.ExecContext(ctx, bet.Id, bet.CustomerId, bet.Status, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
		return err
	}

	insertBetSQL := "INSERT INTO bets(id, customer_id, status, selection_id, selection_coefficient, payment, payout) VALUES (?, ?, ?, ?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.CustomerId, bet.Status, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Payout)
	return err
}

// UpdateBet updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *BetRepository) UpdateBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.Bet) error {
	updateBetSQL := "UPDATE bets SET customer_id=?, status=?, selection_id=?, selection_coefficient=?, payment=?, payout=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.CustomerId, bet.Status, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Payout, bet.Id)
	return err
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}

	return storagemodels.Bet{
		Id:                   id,
		CustomerId:           customerId,
		Status:               status,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
		Payout:               payout,
	}, nil
}
func (r *BetRepository) GetBetsByUserID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetBetsByUserID(ctx, id)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	var domainBets []domainmodels.Bet

	for _, bet := range storageBets {
		domainBet := r.betMapper.MapStorageBetToDomainBet(bet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetsByUserID(ctx context.Context, id string) ([]storagemodels.Bet, error) {
	rows, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id='"+id+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer rows.Close()

	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	var bets []storagemodels.Bet

	for rows.Next() {
		err = rows.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}

		bet := storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		}

		bets = append(bets, bet)
	}

	return bets, nil
}
func (r *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetBetsByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with status"+status)
	}

	var domainBets []domainmodels.Bet

	for _, bet := range storageBets {
		domainBet := r.betMapper.MapStorageBetToDomainBet(bet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetsByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	rows, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status='"+status+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer rows.Close()

	var id string
	var customerId string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	var bets []storagemodels.Bet

	for rows.Next() {
		err = rows.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}

		bet := storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		}

		bets = append(bets, bet)
	}

	return bets, nil
}
