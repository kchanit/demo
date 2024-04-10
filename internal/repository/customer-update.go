package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/request"
)

func (r *repository) UpdateCustomer(ctx context.Context, customerID string, req request.CustomerPutRequest) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET first_name = $1, last_name = $2, image_url = $3, updated_at = $4
	WHERE id = $5
	`, customerTableName)

	_, err := r.pgx.Exec(ctx, query, req.FirstName, req.LastName, req.ImageURL, req.UpdatedAt, customerID)
	if err != nil {
		return err
	}

	return nil
}
