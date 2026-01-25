package configuration

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"math/rand"
	"net/http"
	"robeel-bhatti/go-party-service/internal/constants"
	"strconv"
)

type Middleware struct {
	logger *slog.Logger
	cache  *redis.Client
}

func NewMiddleware(logger *slog.Logger, cache *redis.Client) *Middleware {
	return &Middleware{
		logger: logger,
		cache:  cache,
	}
}

// ValidatePartyId validates the request contains an ID and that the ID is a valid integer.
// It will then set the ID in the request context object to be accessed later in the runtime.
func (m *Middleware) ValidatePartyId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		if id == "" {
			m.logger.Warn("missing party ID in request", "path", r.URL.Path)
			http.Error(w, "party id is required", http.StatusUnprocessableEntity)
			return
		}

		partyId, err := strconv.Atoi(id)

		if err != nil {
			m.logger.Warn("invalid party ID format", "id", id, "error", err)
			http.Error(w, "party id is invalid ", http.StatusUnprocessableEntity)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), constants.PartyIdKey, partyId))
		next(w, r)
	}
}

// Logging logs the HTTP request and what resource was it routed to.
func (m *Middleware) Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g := slog.Group("request", "method", r.Method, "path", r.RequestURI)
		m.logger.Info("http request received", g)
		next(w, r)
	}
}

// Headers sets default outgoing headers all requests will have.
func (m *Middleware) Headers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", constants.ContentType)
		w.Header().Set("Request-ID", strconv.Itoa(rand.Intn(100)))
		next(w, r)
	}
}

// Cache checks if a party exists in the cache before the request is routed to the handler.
// If the party is not found, go to the handler.
func (m *Middleware) Cache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyId, ok := r.Context().Value(constants.PartyIdKey).(int)

		if !ok {
			m.logger.Error("invalid or missing party ID in context")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		m.logger.Info("checking cache for party", "partyID", partyId)
		ck := fmt.Sprintf("%s:%d", constants.ServiceName, partyId)
		res, err := m.cache.Get(r.Context(), ck).Result()

		if errors.Is(err, redis.Nil) {
			m.logger.Info("cache miss for party", "partyID", partyId)
			next(w, r)
			return
		}

		if err != nil {
			m.logger.Error("cache error getting party", "reason", err, "partyId", partyId)
			next(w, r)
			return
		}

		m.logger.Info("cache hit for party", "partyID", partyId)
		_, err = w.Write([]byte(res))

		if err != nil {
			m.logger.Error("error writing response", "reason", err, "partyId", partyId)
			return
		}
	}
}
