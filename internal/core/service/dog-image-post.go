package service

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/internal/core/domain"
)

func (s *service) PostDogImage(ctx context.Context, in domain.DogPostRequest) error {

	dogResponse, err := s.adapter.DogImageGet(string(in.Breed))
	if err != nil {
		return err
	}

	now := s.misc.GetCurrentTimestamp()
	in.CreatedAt = &now
	in.UpdatedAt = &now
	in.ImageURL = dogResponse.Message

	err = s.repo.PostDogImage(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
