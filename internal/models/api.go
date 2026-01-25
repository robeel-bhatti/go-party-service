package models

// PartyCreateDTO is a struct that
// will be used to map the POST request payload to.
type PartyCreateDTO struct {
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	MiddleName  *string           `json:"middle_name"`
	Email       string            `json:"email"`
	PhoneNumber string            `json:"phone_number"`
	Address     *AddressCreateDTO `json:"address"`
	CreatedBy   string            `json:"created_by"`
	UpdatedBy   string            `json:"updated_by"`
}

// AddressCreateDTO is a struct that
// will be used to map the POST request payload to.
type AddressCreateDTO struct {
	StreetOne string  `json:"street_one"`
	StreetTwo *string `json:"street_two"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	ZipCode   string  `json:"zip_code"`
	Country   string  `json:"country"`
	CreatedBy string  `json:"created_by"`
	UpdatedBy string  `json:"updated_by"`
}
