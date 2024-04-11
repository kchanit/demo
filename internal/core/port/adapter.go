package port

import "github.com/centraldigital/cfw-cms-bff/pkg/model/response"

type Adapter interface {
	DogImageGet(breed string) (*response.DogImageReponse, error)
}
