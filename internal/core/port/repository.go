package port

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/internal/repository/entity"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

type Repository interface {
	CreateCustomer(ctx context.Context, req request.CustomerPostRequest) error
	GetAllCustomers(ctx context.Context) ([]entity.Customer, error)
	GetCustomerById(ctx context.Context, customerID string) (*entity.Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customerID string, req request.CustomerPutRequest) error
	DeleteCustomer(ctx context.Context, customerID string) error
}
