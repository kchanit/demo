package handler

import (
	"errors"
	"net/http"

	errorresponse "github.com/centraldigital/cfw-cms-bff/pkg/model/error-response"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerCustomerPut(c *gin.Context, customerID string) {
	req := request.CustomerPutRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.svc.UpdateCustomer(c, customerID, req)
	if err != nil {
		if errors.Is(err, &errorresponse.ErrGetCustomer) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}
