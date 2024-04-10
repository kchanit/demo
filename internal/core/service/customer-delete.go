package service

import (
	"context"

	errorresponse "github.com/centraldigital/cfw-cms-bff/pkg/model/error-response"
)

func (s *service) DeleteCustomer(ctx context.Context, customerID string) error {
	// check if user exists
	customer, _ := s.repo.GetCustomerById(ctx, customerID)
	if customer == nil {
		return &errorresponse.ErrGetCustomer
	}

	err := s.repo.DeleteCustomer(ctx, customerID)
	if err != nil {
		return err
	}

	return nil
}
