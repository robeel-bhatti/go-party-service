package internal

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

// PartyRow struct representing a party row from database
// pointer type means the field is nullable.
type PartyRow struct {
	ID             int       `db:"id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	MiddleName     *string   `db:"middle_name"`
	Email          string    `db:"email"`
	PhoneNumber    string    `db:"phone_number"`
	AddressID      int       `db:"address_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	CreatedBy      string    `db:"created_by"`
	UpdatedBy      string    `db:"updated_by"`
	AddrID         int       `db:"address_pk"`
	AddrStreetOne  string    `db:"street_one"`
	AddrStreetTwo  *string   `db:"street_two"`
	AddrCity       string    `db:"city"`
	AddrState      string    `db:"state"`
	AddrPostalCode string    `db:"postal_code"`
	AddrCountry    string    `db:"country"`
	AddrHash       string    `db:"hash"`
	AddrCreatedAt  time.Time `db:"address_created_at"`
	AddrUpdatedAt  time.Time `db:"address_updated_at"`
	AddrCreatedBy  string    `db:"address_created_by"`
	AddrUpdatedBy  string    `db:"address_updated_by"`
}

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

func (r *PartyRepository) GetById(ctx context.Context) (*PartyRow, error) {
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
			a.postal_code,
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
	row, err := r.db.Query(ctx, q, ctx.Value(partyIdKey))
	if err != nil {
		return nil, err
	}
	defer row.Close()

	partyRow, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[PartyRow])
	if err != nil {
		return nil, err
	}

	return &partyRow, nil
}
