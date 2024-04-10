package repository

import (
	"context"
	"fmt"
)

func (r *repository) DeleteCustomer(ctx context.Context, customerID string) error {
	query := fmt.Sprintf(`
	DELETE FROM %s
	WHERE id = $1
	`, customerTableName)

	_, err := r.pgx.Exec(ctx, query, customerID)
	if err != nil {
		return err
	}

	return nil
}
