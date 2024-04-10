package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

func (r *repository) CreateCustomer(ctx context.Context, req request.CustomerPostRequest) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (first_name, last_name, email, image_url, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`, customerTableName)

	var customerID int64
	err := r.pgx.QueryRow(ctx, query, req.FirstName, req.LastName, req.Email, req.ImageURL, req.CreatedAt, req.UpdatedAt).Scan(&customerID)
	if err != nil {
		return err
	}

	return nil
}
