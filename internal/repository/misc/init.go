package misc

import (
	"context"
	"time"
)

type MiscRepository interface {
	GetCurrentTimestamp() time.Time
}

type MiscRepositoryImpl struct{}

func New(context.Context) (MiscRepository, error) {
	return &MiscRepositoryImpl{}, nil
}
