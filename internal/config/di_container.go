package config

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"robeel-bhatti/go-party-service/internal/controller"
	"robeel-bhatti/go-party-service/internal/repository"
	"robeel-bhatti/go-party-service/internal/service"
)

type Container struct {
	partyController *controller.PartyController
	partyService    *service.PartyService
	partyRepository *repository.PartyRepository
}

func NewContainer(logger *slog.Logger, db *pgxpool.Pool) *Container {
	pr := repository.NewPartyRepository(logger, db)
	ps := service.NewPartyService(logger, pr)
	pc := controller.NewPartyController(logger, ps)
	return &Container{
		partyController: pc,
		partyService:    ps,
		partyRepository: pr,
	}
}
