package handler

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/internal/handler/dto"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

// @Summary      Order credit notation
// @Tags         Order Management
// @Accept       json
// @Security 	 ApiKeyAuth
// @Param		 data body request.DogPostRequest true "body"
// @Produce      json
// @Success      204
// @Failure 	 400  {object} 	errormodel.BusinessError
// @Failure 	 409  {object} 	errormodel.BusinessError
// @Router /v1/dog/{breed} [post]
func (h *handler) HandlerDogImagePost(c context.Context, req request.DogPostRequest) error {
	return h.svc.PostDogImage(c, dto.ToDomainGetDocumentRequest(&req))
}
