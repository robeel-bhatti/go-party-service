package service

import (
	"log/slog"
	"robeel-bhatti/go-party-service/internal/repository"
)

type PartyService struct {
	Log *slog.Logger
	Pr  *repository.PartyRepository
}
