package service

import (
	"robeel-bhatti/go-party-service/internal/models"
)

func mapToPartyResponseDTO(pr *models.PartyReadDTO) *models.PartyResponseDTO {
	return &models.PartyResponseDTO{
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
		Address:     mapToAddressResponseDTO(pr),
	}
}

func mapToAddressResponseDTO(pr *models.PartyReadDTO) *models.AddressResponseDTO {
	return &models.AddressResponseDTO{
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
