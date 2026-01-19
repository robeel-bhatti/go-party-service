package internal

import (
	"errors"
	"net/http"
	"time"
)

func mapToPartyDTO(pr *PartyRow) *Party {
	return &Party{
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

func mapToAddressDTO(pr *PartyRow) *Address {
	return &Address{
		ID:         pr.AddrID,
		StreetOne:  pr.AddrStreetOne,
		StreetTwo:  pr.AddrStreetTwo,
		City:       pr.AddrCity,
		State:      pr.AddrState,
		PostalCode: pr.AddrPostalCode,
		Country:    pr.AddrCountry,
		Hash:       pr.AddrHash,
		CreatedAt:  pr.AddrCreatedAt,
		UpdatedAt:  pr.AddrUpdatedAt,
		CreatedBy:  pr.AddrCreatedBy,
		UpdatedBy:  pr.AddrUpdatedBy,
	}
}

func mapToPartyError(path string, err error) *PartyError {
	pe := &PartyError{
		Timestamp: time.Now(),
		Path:      path,
		Message:   err.Error(),
	}

	for e, c := range ErrMap {
		if errors.Is(err, e) {
			pe.Code = c
			pe.Status = http.StatusText(c)
			break
		}
	}

	return pe
}
