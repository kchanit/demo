package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-cms-bff/internal/repository/entity"
)

func (r *repository) GetAllCustomers(ctx context.Context) ([]entity.Customer, error) {
	query := fmt.Sprintf(`
        SELECT id, first_name, last_name, image_url, created_at, updated_at
        FROM %s
    `, customerTableName)

	rows, err := r.pgx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(
			&customer.CustomerID,
			&customer.FirstName,
			&customer.LastName,
			&customer.ImageURL,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
