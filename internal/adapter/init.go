package adapter

import (
	"errors"
	"net/http"

	domainadapter "github.com/centraldigital/cfw-cms-bff/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-cms-bff/internal/core/port"
	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
)

var ContentTypeJsonHeader = map[string]string{"Content-Type": "application/json"}

type adapter struct {
	httpClient adaptor.ClientAdaptor
}

func New(httpClient *http.Client) port.Adapter {
	return &adapter{
		httpClient: adaptor.NewClient(httpClient),
	}
}

func handleWAppError[T any](readResponse adaptor.ReadResponse[T]) error {
	var clientError domainadapter.WAppClientError
	if err := domainadapter.ParseErrorAsJson(readResponse.Data, &clientError); err != nil {
		return err
	}

	if clientError.Code == "" && clientError.Description == "" {
		return errors.New(string(readResponse.Data))
	}
	return clientError
}
