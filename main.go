package main

import (
	"robeel-bhatti/go-party-service/internal/config"
)

func main() {
	app := &config.App{}
	app.Start()
}
