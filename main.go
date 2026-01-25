package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"robeel-bhatti/go-party-service/internal/configuration"
)

func main() {
	app := configuration.NewApp()
	app.Run(context.Background())
}
