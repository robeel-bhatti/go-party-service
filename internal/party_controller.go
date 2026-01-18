package internal

import (
	"log/slog"
	"net/http"
)

type PartyController struct {
	logger       *slog.Logger
	partyService *PartyService
}

func NewPartyController(logger *slog.Logger, ps *PartyService) *PartyController {
	return &PartyController{
		logger:       logger,
		partyService: ps,
	}
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
