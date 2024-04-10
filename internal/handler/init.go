package handler

import (
	"github.com/centraldigital/cfw-cms-bff/internal/core/port"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandlerCustomerPost(c *gin.Context)
	HandlerCustomerIdGet(c *gin.Context, customerID string)
	HandlerCustomerGetAll(c *gin.Context)
	HandlerCustomerPut(c *gin.Context, customerID string)
	HandlerCustomerDelete(c *gin.Context, customerID string)
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
