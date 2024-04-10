package service

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/response"
)

func (s *service) GetCustomerById(ctx context.Context, customerID string) (*response.CustomerGetByIdResponse, error) {
	customer, err := s.repo.GetCustomerById(ctx, customerID)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	return &response.CustomerGetByIdResponse{
		CustomerID: customer.CustomerID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		ImageURL:   customer.ImageURL,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  customer.UpdatedAt,
	}, nil
}
