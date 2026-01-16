package service

import (
	"log/slog"
	"robeel-bhatti/go-party-service/internal/repository"
)

type PartyService struct {
	logger    *slog.Logger
	partyRepo *repository.PartyRepository
}

func NewPartyService(logger *slog.Logger, pr *repository.PartyRepository) *PartyService {
	return &PartyService{
		logger:    logger,
		partyRepo: pr,
	}
}
