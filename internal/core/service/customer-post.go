package service

import (
	"context"
	"fmt"

	errorresponse "github.com/centraldigital/cfw-cms-bff/pkg/model/error-response"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

func (s *service) CreateCustomer(ctx context.Context, req request.CustomerPostRequest) error {
	now := s.misc.GetCurrentTimestamp()
	req.CreatedAt = &now
	req.UpdatedAt = &now

	// check if email already exists
	customer, err := s.repo.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if customer != nil {
		fmt.Println("email already exists")
		fmt.Println(customer)
		return &errorresponse.ErrDuplicateCustomerEmail
	}

	// create customer
	err = s.repo.CreateCustomer(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
