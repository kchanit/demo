package service

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/response"
)

func (s *service) GetAllCustomers(ctx context.Context) (*response.CustomerGetAllResponse, error) {
	customers, err := s.repo.GetAllCustomers(ctx)
	if err != nil {
		return nil, err
	}

	if customers == nil {
		return nil, nil
	}

	return &response.CustomerGetAllResponse{
		Customers: customers,
	}, nil
}
