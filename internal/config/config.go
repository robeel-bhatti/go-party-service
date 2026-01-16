package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"os"
	"robeel-bhatti/go-party-service/internal/controller"
)

type App struct {
	logger *slog.Logger
}

// NewApp creates and returns an instance of an App struct that represents
// the main application object
func NewApp(logger *slog.Logger) *App {
	return &App{
		logger: logger,
	}
}

func (app *App) Run(ctx context.Context) {
	db, err := app.connectToDB(ctx)
	defer db.Close()

	if err != nil {
		app.logger.Error("failed to connect to database", "reason", err)
		os.Exit(1)
	}

	_, err = app.connectToCache(ctx) // TODO: add back cache variable when ready to use
	if err != nil {
		app.logger.Error("failed to connect to cache", "reason", err)
		os.Exit(1)
	}

	c := NewContainer(app.logger, db)
	mux := app.getMultiplexer(c.partyController)
	err = http.ListenAndServe(os.Getenv("PORT"), mux)

	if err != nil {
		app.logger.Error("failed to start server", "reason", err)
		os.Exit(1)
	}
}

func (app *App) connectToDB(ctx context.Context) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *App) connectToCache(ctx context.Context) (*redis.Client, error) {
	opts := &redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	}
	cache := redis.NewClient(opts)

	_, err := cache.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (app *App) getMultiplexer(pc *controller.PartyController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parties/{id}", pc.GetPartyById)
	mux.HandleFunc("PATCH /parties/{id}", pc.UpdateParty)
	mux.HandleFunc("DELETE /parties/{id}", pc.DeleteParty)
	mux.HandleFunc("POST /parties", pc.CreateParty)
	mux.HandleFunc("GET /parties", pc.GetParties)
	return mux
}
