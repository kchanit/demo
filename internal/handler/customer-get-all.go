package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerCustomerGetAll(c *gin.Context) {
	resp, err := h.svc.GetAllCustomers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if resp == nil {
		c.Status(http.StatusNoContent)
	}

	c.JSON(http.StatusOK, resp)
}
