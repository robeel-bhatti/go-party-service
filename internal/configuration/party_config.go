package configuration

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"os"
	"robeel-bhatti/go-party-service/internal/constants"
)

type App struct{}

// NewApp creates and returns an instance of an App struct that represents
// the main application object.
func NewApp() *App {
	return &App{}
}

// Run starts up the application and creates all the necessary
// configuration the app needs in order to run.
func (app *App) Run(ctx context.Context) {
	logger := app.configureLogger()

	db, err := app.configureDB(ctx)
	defer db.Close()

	if err != nil {
		logger.Error("failed to connect to database", "reason", err)
		os.Exit(1)
	}

	cache, err := app.configureCache(ctx)
	if err != nil {
		logger.Error("failed to connect to cache", "reason", err)
		os.Exit(1)
	}

	c := NewContainer(logger, db, cache)
	mux := app.configureMultiplexer(c)
	err = http.ListenAndServe(os.Getenv("PORT"), mux)

	if err != nil {
		logger.Error("failed to start server", "reason", err)
		os.Exit(1)
	}
}

// configureLogger creates the app-wide logger.
func (app *App) configureLogger() *slog.Logger {
	var logLevel slog.Level
	value := os.Getenv("LOG_LEVEL")

	switch value {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: logLevel}
	return slog.New(slog.NewJSONHandler(os.Stderr, opts)).With("service", constants.ContentType)
}

// configureDB creates a connection to the database
// and returns a database object that can be used to query.
func (app *App) configureDB(ctx context.Context) (*pgxpool.Pool, error) {
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

// configureCache creates a connection to the cache
// and returns a cache object that can be used to query.
func (app *App) configureCache(ctx context.Context) (*redis.Client, error) {
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

// configureMultiplexer creates and returns this application's main multiplexer
// so that incoming HTTP requests can be routed to the correct endpoint.
func (app *App) configureMultiplexer(c *Container) *http.ServeMux {
	mux := http.NewServeMux()
	mw := c.middleware
	pc := c.partyController
	mux.HandleFunc("GET /api/v1/parties/{id}", mw.ValidatePartyId(mw.Logging(mw.Headers(mw.Cache(pc.GetPartyById)))))
	mux.HandleFunc("PATCH /api/v1/parties/{id}", c.partyController.UpdateParty)
	mux.HandleFunc("POST /api/v1/parties", c.partyController.CreateParty)
	return mux
}
