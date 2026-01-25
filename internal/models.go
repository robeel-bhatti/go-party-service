package internal

import "time"

// PartyRow represents a row from the database.
// Pointer type means the field is nullable.
type PartyRow struct {
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

// PartyDTO domain object
type PartyDTO struct {
	ID          int         `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	MiddleName  *string     `json:"middle_name"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phone_number"`
	Address     *AddressDTO `json:"address"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CreatedBy   string      `json:"created_by"`
	UpdatedBy   string      `json:"updated_by"`
}

// AddressDTO domain object
type AddressDTO struct {
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
