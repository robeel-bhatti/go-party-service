package internal

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type PartyService struct {
	logger    *slog.Logger
	partyRepo *PartyRepository
	cache     *redis.Client
}

func NewPartyService(logger *slog.Logger, pr *PartyRepository, ca *redis.Client) *PartyService {
	return &PartyService{
		logger:    logger,
		partyRepo: pr,
		cache:     ca,
	}
}
