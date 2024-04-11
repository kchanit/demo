package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-cms-bff/internal/core/domain"
)

func (r *repository) PostDogImage(ctx context.Context, request domain.DogPostRequest) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (breed, image_url, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`, dogTableName)

	var dogID int64
	err := r.pgx.QueryRow(ctx, query, request.Breed, request.ImageURL, request.CreatedAt, request.UpdatedAt).Scan(&dogID)
	if err != nil {
		return err
	}

	return nil
}
