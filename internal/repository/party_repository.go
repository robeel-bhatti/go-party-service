package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
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
