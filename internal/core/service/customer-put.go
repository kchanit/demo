package service

import (
	"context"

	errorresponse "github.com/centraldigital/cfw-cms-bff/pkg/model/error-response"
	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

func (s *service) UpdateCustomer(ctx context.Context, customerID string, req request.CustomerPutRequest) error {
	now := s.misc.GetCurrentTimestamp()
	req.UpdatedAt = &now

	// check if user exists
	customer, _ := s.repo.GetCustomerById(ctx, customerID)
	if customer == nil {
		return &errorresponse.ErrGetCustomer
	}

	err := s.repo.UpdateCustomer(ctx, customerID, req)
	if err != nil {
		return err
	}
	return nil
}
