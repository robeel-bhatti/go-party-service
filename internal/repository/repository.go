package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"robeel-bhatti/go-party-service/internal/models"
)

type PartyRepository struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewPartyRepository(logger *slog.Logger, db *pgxpool.Pool) *PartyRepository {
	return &PartyRepository{
		logger: logger,
		db:     db,
	}
}

// GetById checks the database source to get an existing party with the provided ID
// Only one row has to be returned otherwise return an error.
func (r *PartyRepository) GetById(ctx context.Context, partyId int) (*models.PartyReadDTO, error) {
	q := `
		SELECT
			p.id,
			p.first_name,
			p.last_name, 
			p.middle_name,
			p.email,
			p.phone_number,
			p.address_id,
			p.created_at,
			p.updated_at,
			p.created_by,
			p.updated_by,
			a.id AS address_pk,
			a.street_one,
			a.street_two,
			a.city,
			a.state,
			a.zip_code,
			a.country,
			a.hash,
			a.created_at AS address_created_at,
			a.updated_at AS address_updated_at,
			a.created_by AS address_created_by,
			a.updated_by AS address_updated_by
		FROM party_service.party p
		LEFT JOIN party_service.address a ON p.address_id = a.id
		WHERE p.id = $1;
	`
	row, err := r.db.Query(ctx, q, partyId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	partyRow, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.PartyReadDTO])
	if err != nil {
		return nil, err
	}
	return &partyRow, nil
}
