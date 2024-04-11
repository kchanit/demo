package request

import "github.com/centraldigital/cfw-cms-bff/pkg/model/enum"

type DogPostRequest struct {
	Breed enum.DogBreed `json:"breed" uri:"breed" example:"hound"`
}
