package handler

import (
	"fmt"
	"net/http"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
	"github.com/gin-gonic/gin"
)

// @Summary      Order credit notation
// @Tags         Order Management
// @Accept       json
// @Security 	 ApiKeyAuth
// @Param		 data body request.CustomerPostRequest true "body"
// @Produce      json
// @Success      201  {object}  response.CustomerCreateResponse
// @Failure 	 400  {object} 	errormodel.BusinessError
// @Failure 	 409  {object} 	errormodel.BusinessError
// @Router /v1/customer [post]
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
