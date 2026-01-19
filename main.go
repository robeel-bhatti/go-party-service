package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"robeel-bhatti/go-party-service/internal"
)

func main() {
	app := internal.NewApp()
	app.Run(context.Background())
}
