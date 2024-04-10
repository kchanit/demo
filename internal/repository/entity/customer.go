package entity

import "time"

type Customer struct {
	CustomerID string     `json:"customer_id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	ImageURL   string     `json:"image_url"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
