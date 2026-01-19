package internal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type (
	Container struct {
		db              *pgxpool.Pool
		cache           *redis.Client
		middleware      *Middleware
		partyController *PartyController
		partyService    *PartyService
		partyRepository *PartyRepository
	}
)

func NewContainer(logger *slog.Logger, db *pgxpool.Pool, cache *redis.Client) *Container {
	mw := NewMiddleware(logger, cache)
	pr := NewPartyRepository(logger, db)
	ps := NewPartyService(logger, pr, cache)
	pc := NewPartyController(logger, ps)
	return &Container{
		db:              db,
		cache:           cache,
		middleware:      mw,
		partyController: pc,
		partyService:    ps,
		partyRepository: pr,
	}
}
