package response

import (
	"time"

	"github.com/centraldigital/cfw-cms-bff/internal/repository/entity"
)

type CustomerCreateResponse struct {
	CustomerID string `json:"customerID"`
}

type CustomerGetByIdResponse struct {
	CustomerID string     `json:"customer_id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	ImageURL   string     `json:"image_url"`
	Email      string     `json:"email"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type CustomerGetAllResponse struct {
	Customers []entity.Customer `json:"customers"`
}
