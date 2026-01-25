package models

import "time"

// PartyReadDTO is a struct that represents
// an entire Party entity read from the database.
// This struct is only used for reads.
type PartyReadDTO struct {
	ID            int       `db:"id"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	MiddleName    *string   `db:"middle_name"`
	Email         string    `db:"email"`
	PhoneNumber   string    `db:"phone_number"`
	AddressID     int       `db:"address_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatedBy     string    `db:"created_by"`
	UpdatedBy     string    `db:"updated_by"`
	AddrID        int       `db:"address_pk"`
	AddrStreetOne string    `db:"street_one"`
	AddrStreetTwo *string   `db:"street_two"`
	AddrCity      string    `db:"city"`
	AddrState     string    `db:"state"`
	AddrZipCode   string    `db:"zip_code"`
	AddrCountry   string    `db:"country"`
	AddrHash      string    `db:"hash"`
	AddrCreatedAt time.Time `db:"address_created_at"`
	AddrUpdatedAt time.Time `db:"address_updated_at"`
	AddrCreatedBy string    `db:"address_created_by"`
	AddrUpdatedBy string    `db:"address_updated_by"`
}

// PartyWriteDTO is a struct whose values
// will be written to the database to the party table.
type PartyWriteDTO struct {
	FirstName   string  `db:"first_name"`
	LastName    string  `db:"last_name"`
	MiddleName  *string `db:"middle_name"`
	Email       string  `db:"email"`
	PhoneNumber string  `db:"phone_number"`
	AddressID   int     `db:"address_id"`
	CreatedBy   string  `db:"created_by"`
	UpdatedBy   string  `db:"updated_by"`
}

// AddressWriteDTO is a struct whose values
// will be written to the database to the address table.
type AddressWriteDTO struct {
	AddrStreetOne string  `db:"street_one"`
	AddrStreetTwo *string `db:"street_two"`
	AddrCity      string  `db:"city"`
	AddrState     string  `db:"state"`
	AddrZipCode   string  `db:"zip_code"`
	AddrCountry   string  `db:"country"`
	AddrHash      string  `db:"hash"`
	AddrCreatedBy string  `db:"address_created_by"`
	AddrUpdatedBy string  `db:"address_updated_by"`
}
