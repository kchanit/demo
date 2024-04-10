package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/centraldigital/cfw-cms-bff/internal/repository/entity"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetCustomerById(ctx context.Context, customerID string) (*entity.Customer, error) {
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name, image_url, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, customerTableName)

	var customer entity.Customer
	err := r.pgx.QueryRow(ctx, query, customerID).Scan(
		&customer.CustomerID,
		&customer.FirstName,
		&customer.LastName,
		&customer.ImageURL,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}
