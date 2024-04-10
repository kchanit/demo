package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerCustomerIdGet(c *gin.Context, customerID string) {
	resp, err := h.svc.GetCustomerById(c, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if resp == nil {
		c.Status(http.StatusNoContent)
	}

	c.JSON(http.StatusOK, resp)
}
