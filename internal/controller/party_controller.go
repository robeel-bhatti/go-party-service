package controller

import (
	"log/slog"
	"net/http"
	"robeel-bhatti/go-party-service/internal/service"
)

type PartyController struct {
	logger       *slog.Logger
	partyService *service.PartyService
}

func NewPartyController(logger *slog.Logger, ps *service.PartyService) *PartyController {
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
