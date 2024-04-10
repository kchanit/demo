package handler

import (
	"errors"
	"net/http"

	errorresponse "github.com/centraldigital/cfw-cms-bff/pkg/model/error-response"
	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerCustomerDelete(c *gin.Context, customerID string) {
	err := h.svc.DeleteCustomer(c, customerID)
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
