package handler

import (
	"fmt"
	"net/http"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerCustomerPost(c *gin.Context) {
	req := request.CustomerPostRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.svc.CreateCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusCreated)
}
