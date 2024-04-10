package domainadapter

import (
	"encoding/json"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
)

type WAppClientError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (da WAppClientError) Error() string {
	return da.Description
}

func ParseErrorAsJson[T error](data []byte, errorJson *T) error {
	err := json.Unmarshal(data, errorJson)
	if err != nil {
		return errormodel.ClientErrorDefaultCode(
			fmt.Sprintf("error response: %s -- unable to unmarhal json: %v", data, err),
		)
	}
	return nil
}
