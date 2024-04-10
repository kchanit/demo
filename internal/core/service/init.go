package service

import (
	"github.com/centraldigital/cfw-cms-bff/internal/core/port"
	"github.com/centraldigital/cfw-cms-bff/internal/repository/misc"
)

type service struct {
	adapter port.Adapter
	repo    port.Repository
	misc    misc.MiscRepository
}

func New(adapter port.Adapter, repo port.Repository, misc misc.MiscRepository) port.Service {
	return &service{
		adapter: adapter,
		repo:    repo,
		misc:    misc,
	}
}
