package internal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	logger *slog.Logger
}

// NewApp creates and returns an instance of an App struct that represents
// the main application object/
func NewApp(logger *slog.Logger) *App {
	return &App{
		logger: logger,
	}
}

// Run starts up the application and creates all the necessary
// configuration the app needs in order to run.
func (app *App) Run(ctx context.Context) {
	db, err := app.connectToDB(ctx)
	defer db.Close()

	if err != nil {
		app.logger.Error("failed to connect to database", "reason", err)
		os.Exit(1)
	}

	cache, err := app.connectToCache(ctx)
	if err != nil {
		app.logger.Error("failed to connect to cache", "reason", err)
		os.Exit(1)
	}

	c := NewContainer(app.logger, db, cache)
	mux := app.getMultiplexer(ctx, c)
	err = http.ListenAndServe(os.Getenv("PORT"), mux)

	if err != nil {
		app.logger.Error("failed to start server", "reason", err)
		os.Exit(1)
	}
}

// connectToDB creates a connection to the database
// and returns a database object that can be used to query.
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

// connectToCache creates a connection to the cache
// and returns a cache object that can be used to query.
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

// getMultiplexer creates and returns this application's main multiplexer
// so that incoming HTTP requests can be routed to the correct endpoint.
func (app *App) getMultiplexer(ctx context.Context, c *Container) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parties/{id}", CacheMiddleware(ctx, c.partyController.GetPartyById, c.cache))
	mux.HandleFunc("PATCH /parties/{id}", c.partyController.UpdateParty)
	mux.HandleFunc("DELETE /parties/{id}", c.partyController.DeleteParty)
	mux.HandleFunc("POST /parties", c.partyController.CreateParty)
	mux.HandleFunc("GET /parties", c.partyController.GetParties)
	return mux
}
