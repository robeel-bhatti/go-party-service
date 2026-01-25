package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"robeel-bhatti/go-party-service/internal"
	"robeel-bhatti/go-party-service/internal/controller"
	"robeel-bhatti/go-party-service/internal/repository"
)

// PartyService handles business logic for parties
type PartyService struct {
	logger    *slog.Logger
	partyRepo *repository.PartyRepository
	cache     *redis.Client
}

func NewPartyService(logger *slog.Logger, pr *repository.PartyRepository, ca *redis.Client) *PartyService {
	return &PartyService{
		logger:    logger,
		partyRepo: pr,
		cache:     ca,
	}
}

// GetPartyById performs the biz logic to get a party entity from the data layer
// and return a party domain object to API layer.
// Set the party in the cache afterward, since if this method is hit there was a party cache miss.
func (s *PartyService) GetPartyById(ctx context.Context, partyId int) (*controller.PartyDTO, error) {
	pr, err := s.partyRepo.GetById(ctx, partyId)

	if err != nil {
		s.logger.Error("database error", "reason", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("failed to get party %d: %w", partyId, controller.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get party %d: %w", partyId, controller.ErrInternalServerError)
	}

	res := internal.mapToPartyDTO(pr)
	s.setInCache(ctx, partyId, res)
	return res, nil
}

// setInCache sets the provided party in the cache.
// if an error occurs, either marshalling or setting in redis, log the error and continue
// we don't want to return an exception to the invoking client in this case.
func (s *PartyService) setInCache(ctx context.Context, partyId int, party *controller.PartyDTO) {
	b, err := json.Marshal(party)
	if err != nil {
		s.logger.Error("error marshalling party to set in cache", "reason", err)
		return
	}

	ck := fmt.Sprintf("%s:%d", internal.serviceName, partyId)
	_, err = s.cache.Set(ctx, ck, b, 0).Result()
	if err != nil {
		s.logger.Error("error setting in cache", "reason", err)
	}
}
