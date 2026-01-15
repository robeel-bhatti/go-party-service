package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

type App struct {
	log      *slog.Logger
	httpPort int
	dbPort   int
	dbHost   string
	dbName   string
	dbUser   string
	dbPass   string
}

// NewApp creates and returns an instance of an App struct
// IRL would pull these values from env vars etc.
func NewApp(log *slog.Logger) *App {
	return &App{
		log:      log,
		httpPort: 8080,
		dbPort:   5432,
		dbHost:   "localhost",
		dbName:   "Party",
		dbUser:   "admin",
		dbPass:   "admin",
	}
}

func (app *App) Run() {
	c := NewContainer()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parties/{id}", c.Pc.GetPartyById)
	mux.HandleFunc("PATCH /parties/{id}", c.Pc.UpdateParty)
	mux.HandleFunc("DELETE /parties/{id}", c.Pc.DeleteParty)
	mux.HandleFunc("POST /parties", c.Pc.CreateParty)
	mux.HandleFunc("GET /parties", c.Pc.GetParties)
}

func (app *App) connectToDB() *sql.DB {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		app.dbHost, app.dbPort, app.dbUser, app.dbPass, app.dbName,
	)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
