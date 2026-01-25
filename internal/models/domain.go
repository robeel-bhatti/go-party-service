package models

import "time"

// PartyResponseDTO domain object
type PartyResponseDTO struct {
	ID          int                 `json:"id"`
	FirstName   string              `json:"first_name"`
	LastName    string              `json:"last_name"`
	MiddleName  *string             `json:"middle_name"`
	Email       string              `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	Address     *AddressResponseDTO `json:"address"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	CreatedBy   string              `json:"created_by"`
	UpdatedBy   string              `json:"updated_by"`
}

// AddressResponseDTO domain object
type AddressResponseDTO struct {
	ID        int       `json:"id"`
	StreetOne string    `json:"street_one"`
	StreetTwo *string   `json:"street_two"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Hash      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}
