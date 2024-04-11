package dto

import (
	"github.com/centraldigital/cfw-cms-bff/internal/core/domain"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

func ToDomainGetDocumentRequest(d *request.DogPostRequest) domain.DogPostRequest {
	return domain.DogPostRequest{
		Breed: d.Breed,
	}
}

// func FromDomainGetDocumentResponse(d *domain.DogPostRequest) *response.DogPostResponse {
// 	return &response.DogPostResponse{}
// }
