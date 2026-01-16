package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"log/slog"
	"os"
	"robeel-bhatti/go-party-service/internal/config"
)

func main() {
	logger := configureLogger()
	app := config.NewApp(logger)
	app.Run(context.Background())
}

func configureLogger() *slog.Logger {
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
	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).With("service", os.Getenv("SERVICE_NAME"))
}
