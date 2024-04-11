package handler

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/internal/core/port"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandlerCustomerPost(c *gin.Context)
	HandlerCustomerIdGet(c *gin.Context, customerID string)
	HandlerCustomerGetAll(c *gin.Context)
	HandlerCustomerPut(c *gin.Context, customerID string)
	HandlerCustomerDelete(c *gin.Context, customerID string)

	HandlerDogImagePost(c context.Context, request request.DogPostRequest) error
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
