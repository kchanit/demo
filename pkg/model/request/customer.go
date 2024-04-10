package request

import "time"

type CustomerPostRequest struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	ImageURL  string     `json:"image_url"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CustomerGetByIdRequest struct {
	CustomerId string
}

type CustomerPutRequest struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	ImageURL  string     `json:"image_url"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CustomerDeleteRequest struct {
	CustomerId string
}
