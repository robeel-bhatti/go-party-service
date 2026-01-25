package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"robeel-bhatti/go-party-service/internal"
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

// GetPartyById is the application endpoint that gets a party entity via the provided ID in the HTTP request.
// if any error occur marshalling or writing the response to the buffer, return a default plaintext 500 response as
// a final measure.
func (pc *PartyController) GetPartyById(w http.ResponseWriter, r *http.Request) {
	partyId := r.Context().Value(internal.partyIdKey).(int)
	pc.logger.Info("controller retrieved request to get party from database", "partyId", partyId)
	res, err := pc.partyService.GetPartyById(r.Context(), partyId)

	if err != nil {
		pe := service.mapToPartyError(r.RequestURI, err)
		b, err := json.Marshal(pe)
		if err != nil {
			http.Error(w, "error marshalling response", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(pe.Code)
		_, err = w.Write(b)
		if err != nil {
			http.Error(w, "error writing response", http.StatusInternalServerError)
			return
		}
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return
	}
	return
}

func (pc *PartyController) UpdateParty(w http.ResponseWriter, r *http.Request) {

}

func (pc *PartyController) CreateParty(w http.ResponseWriter, r *http.Request) {

}
