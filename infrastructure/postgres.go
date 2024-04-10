package infrastructure

import (
	"context"

	"github.com/centraldigital/cfw-cms-bff/property"
	"github.com/centraldigital/cfw-core-lib/pkg/util/infrastructureutil"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, *pgxscan.API) {
	pgx, scanapi, err := infrastructureutil.NewPostgresWithScanApi(ctx, property.Get().Postgres.PostgresConfig)
	if err != nil {
		panic(err)
	}
	return pgx, scanapi
}
