package config

import (
	"log/slog"
	"robeel-bhatti/go-party-service/internal/controller"
	"robeel-bhatti/go-party-service/internal/repository"
	"robeel-bhatti/go-party-service/internal/service"
)

type Container struct {
	Pc *controller.PartyController
	Ps *service.PartyService
	Pr *repository.PartyRepository
}

func NewContainer() *Container {
	log := slog.Default()
	pr := &repository.PartyRepository{Log: log}
	ps := &service.PartyService{Log: log, Pr: pr}
	pc := &controller.PartyController{Log: log, Ps: ps}
	return &Container{Pc: pc, Ps: ps, Pr: pr}
}
