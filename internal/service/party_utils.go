package service

import (
	"errors"
	"net/http"
	"robeel-bhatti/go-party-service/internal/controller"
	"time"
)

func mapToPartyDTO(pr *controller.PartyReadDTO) *controller.PartyDTO {
	return &controller.PartyDTO{
		ID:          pr.ID,
		FirstName:   pr.FirstName,
		LastName:    pr.LastName,
		MiddleName:  pr.MiddleName,
		Email:       pr.Email,
		PhoneNumber: pr.PhoneNumber,
		CreatedAt:   pr.CreatedAt,
		UpdatedAt:   pr.UpdatedAt,
		CreatedBy:   pr.CreatedBy,
		UpdatedBy:   pr.UpdatedBy,
		Address:     mapToAddressDTO(pr),
	}
}

func mapToAddressDTO(pr *controller.PartyReadDTO) *controller.AddressDTO {
	return &controller.AddressDTO{
		ID:        pr.AddrID,
		StreetOne: pr.AddrStreetOne,
		StreetTwo: pr.AddrStreetTwo,
		City:      pr.AddrCity,
		State:     pr.AddrState,
		ZipCode:   pr.AddrZipCode,
		Country:   pr.AddrCountry,
		Hash:      pr.AddrHash,
		CreatedAt: pr.AddrCreatedAt,
		UpdatedAt: pr.AddrUpdatedAt,
		CreatedBy: pr.AddrCreatedBy,
		UpdatedBy: pr.AddrUpdatedBy,
	}
}

func mapToPartyError(path string, err error) *controller.PartyError {
	pe := &controller.PartyError{
		Timestamp: time.Now(),
		Path:      path,
		Message:   err.Error(),
	}

	for e, c := range controller.ErrMap {
		if errors.Is(err, e) {
			pe.Code = c
			pe.Status = http.StatusText(c)
			break
		}
	}

	return pe
}
