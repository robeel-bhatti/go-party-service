package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
	"time"
)

// PartyService handles business logic for parties
type PartyService struct {
	logger    *slog.Logger
	partyRepo *PartyRepository
	cache     *redis.Client
}

type Party struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  *string   `json:"middle_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     *Address  `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

// Address represents a physical address
type Address struct {
	ID         int       `json:"id"`
	StreetOne  string    `json:"street_one"`
	StreetTwo  *string   `json:"street_two"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	Hash       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
}

func NewPartyService(logger *slog.Logger, pr *PartyRepository, ca *redis.Client) *PartyService {
	return &PartyService{
		logger:    logger,
		partyRepo: pr,
		cache:     ca,
	}
}

// GetPartyById performs the biz logic to get a party entity from the data layer
// and return a party domain object to API layer.
// Set the party in the cache afterward, since if this method is hit there was a party cache miss.
func (s *PartyService) GetPartyById(ctx context.Context, partyId int) (*Party, error) {
	pr, err := s.partyRepo.GetById(ctx, partyId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("party %d %w", partyId, ErrNotFound)
		}
		return nil, fmt.Errorf("party %d %w", partyId, ErrInternalServerError)
	}

	res := mapToPartyDTO(pr)
	s.setInCache(ctx, partyId, res)
	return res, nil
}

// setInCache sets the provided party in the cache.
// if an error occurs, either marshalling or setting in redis, log the error and continue
// we don't want to return an exception to the invoking client in this case.
func (s *PartyService) setInCache(ctx context.Context, partyId int, party *Party) {
	b, err := json.Marshal(party)
	if err != nil {
		s.logger.Error("error marshalling party to set in cache", "reason", err)
		return
	}

	ck := fmt.Sprintf("%s:%d", os.Getenv("SERVICE_NAME"), partyId)
	_, err = s.cache.Set(ctx, ck, b, 0).Result()
	if err != nil {
		s.logger.Error("error setting in cache", "reason", err)
	}
}
