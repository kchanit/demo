package port

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/response"
)

type Service interface {
	CreateCustomer(ctx context.Context, req request.CustomerPostRequest) error
	GetCustomerById(ctx context.Context, customerID string) (*response.CustomerGetByIdResponse, error)
	GetAllCustomers(ctx context.Context) (*response.CustomerGetAllResponse, error)
	UpdateCustomer(ctx context.Context, customerID string, req request.CustomerPutRequest) error
	DeleteCustomer(ctx context.Context, customerID string) error
}
