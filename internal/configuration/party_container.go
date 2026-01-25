package configuration

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"robeel-bhatti/go-party-service/internal/controller"
	"robeel-bhatti/go-party-service/internal/repository"
	"robeel-bhatti/go-party-service/internal/service"
)

type (
	Container struct {
		db              *pgxpool.Pool
		cache           *redis.Client
		middleware      *Middleware
		partyController *controller.PartyController
		partyService    *service.PartyService
		partyRepository *repository.PartyRepository
	}
)

func NewContainer(logger *slog.Logger, db *pgxpool.Pool, cache *redis.Client) *Container {
	mw := NewMiddleware(logger, cache)
	pr := repository.NewPartyRepository(logger, db)
	ps := service.NewPartyService(logger, pr, cache)
	pc := controller.NewPartyController(logger, ps)
	return &Container{
		db:              db,
		cache:           cache,
		middleware:      mw,
		partyController: pc,
		partyService:    ps,
		partyRepository: pr,
	}
}
