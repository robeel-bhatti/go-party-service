package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "Party"
)

type App struct {
}

func (app *App) Start() {
	c := NewContainer()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /parties/{id}", c.Pc.GetPartyById)
	mux.HandleFunc("PATCH /parties/{id}", c.Pc.UpdateParty)
	mux.HandleFunc("DELETE /parties/{id}", c.Pc.DeleteParty)
	mux.HandleFunc("POST /parties", c.Pc.CreateParty)
	mux.HandleFunc("GET /parties", c.Pc.GetParties)
}

func ConnectToDB() {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
}

func ConnectToCache() {

}
