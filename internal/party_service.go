package internal

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type (
	Party        struct{}
	PartyService struct {
		logger    *slog.Logger
		partyRepo *PartyRepository
		cache     *redis.Client
	}
)

func NewPartyService(logger *slog.Logger, pr *PartyRepository, ca *redis.Client) *PartyService {
	return &PartyService{
		logger:    logger,
		partyRepo: pr,
		cache:     ca,
	}
}

func (s *PartyService) GetPartyById(ctx context.Context) (*Party, error) {
	_, err := s.partyRepo.GetById(ctx)
	if err != nil {
		return nil, err
	}
	return &Party{}, nil
}
