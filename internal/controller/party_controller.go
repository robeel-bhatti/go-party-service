package controller

import (
	"log/slog"
	"net/http"
	"robeel-bhatti/go-party-service/internal/service"
)

type PartyController struct {
	Log *slog.Logger
	Ps  *service.PartyService
}

func (pc *PartyController) GetPartyById(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) UpdateParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) DeleteParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) CreateParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) GetParties(w http.ResponseWriter, r *http.Request) {

}
