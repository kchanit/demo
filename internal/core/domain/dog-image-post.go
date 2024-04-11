package domain

import (
	"time"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/enum"
)

type DogPostRequest struct {
	Breed     enum.DogBreed `json:"breed"`
	ImageURL  string        `json:"image_url"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
}
