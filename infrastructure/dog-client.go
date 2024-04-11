package infrastructure

import (
	"context"
	"net/http"
	"time"
)

func NewDogHttpClient(ctx context.Context) *http.Client {
	http := &http.Client{Timeout: 30 * time.Second}
	return http
}
