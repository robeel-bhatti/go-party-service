package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type PartyIdContextKey string

type Middleware struct {
	logger *slog.Logger
	cache  *redis.Client
}

const (
	partyIdKey PartyIdContextKey = "partyId"
)

func NewMiddleware(logger *slog.Logger, cache *redis.Client) *Middleware {
	return &Middleware{
		logger: logger,
		cache:  cache,
	}
}

func (m *Middleware) UpdateContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyId, _ := strconv.Atoi(r.PathValue("id"))
		r = r.WithContext(context.WithValue(r.Context(), partyIdKey, partyId))
		next(w, r)
	}
}

func (m *Middleware) Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g := slog.Group("request", "method", r.Method, "path", r.Context().Value(partyIdKey))
		m.logger.Info("http request received", g)
		next(w, r)
	}
}

func (m *Middleware) Headers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Request-ID", strconv.Itoa(rand.Intn(100)))
		next(w, r)
	}
}

func (m *Middleware) Cache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyId := r.Context().Value(partyIdKey).(int)
		m.logger.Info("checking cache for party", "partyID", partyId)

		ck := fmt.Sprintf("%s:%d", os.Getenv("SERVICE_NAME"), partyId)
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
