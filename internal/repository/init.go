package repository

import (
	"github.com/centraldigital/cfw-cms-bff/internal/core/port"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pgx     *pgxpool.Pool
	scanApi *pgxscan.API
}

func New(pgx *pgxpool.Pool, scanApi *pgxscan.API) port.Repository {
	return &repository{
		pgx:     pgx,
		scanApi: scanApi,
	}
}
