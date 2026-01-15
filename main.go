package main

import (
	"log/slog"
	"os"
	"robeel-bhatti/go-party-service/internal/config"
)

func main() {
	logger := configureLogger()
	app := config.NewApp(logger)
	app.Run()
}

func configureLogger() *slog.Logger {
	var level slog.Level
	val := os.Getenv("LOG_LEVEL")

	switch val {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).With("service", os.Getenv("SERVICE_NAME"))
}
