package internal

import (
	"encoding/json"
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
	res, err := pc.partyService.GetPartyById(r.Context())
	if err != nil {
		pe := mapToPartyError(r.URL.Path, err)
		pc.logger.Info("this is", "message", pe)
		b, err := json.Marshal(pe)
		if err != nil {

		}
		_, err = w.Write(b)
		if err != nil {

		}
		return
	}

	b, err := json.Marshal(res)
	if err != nil {

	}

	_, err = w.Write(b)
	if err != nil {

	}
}

func (pc *PartyController) UpdateParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) DeleteParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) CreateParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) GetParties(w http.ResponseWriter, r *http.Request) {

}
