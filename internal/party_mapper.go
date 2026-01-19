package internal

import (
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

type PartyError struct {
	timestamp time.Time
	path      string
	status    string
	code      int
	message   string
}

func mapToPartyError(p, s string, c int, m string) *PartyError {
	return &PartyError{
		timestamp: time.Now(),
		path:      p,
		status:    s,
		code:      c,
		message:   m,
	}
}
