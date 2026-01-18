package internal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Container struct {
	db              *pgxpool.Pool
	cache           *redis.Client
	partyController *PartyController
	partyService    *PartyService
	partyRepository *PartyRepository
}

func NewContainer(logger *slog.Logger, db *pgxpool.Pool, ca *redis.Client) *Container {
	pr := NewPartyRepository(logger, db)
	ps := NewPartyService(logger, pr, ca)
	pc := NewPartyController(logger, ps)
	return &Container{
		db:              db,
		cache:           ca,
		partyController: pc,
		partyService:    ps,
		partyRepository: pr,
	}
}
