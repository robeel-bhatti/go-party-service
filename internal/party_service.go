package internal

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
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

func (s *PartyService) GetPartyById(ctx context.Context) (*Party, error) {
	pr, err := s.partyRepo.GetById(ctx)
	if err != nil {
		return nil, err
	}
	return mapToPartyDTO(pr), nil
}
