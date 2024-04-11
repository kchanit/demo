package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/centraldigital/cfw-cms-bff/pkg/model/response"
)

func (a *adapter) DogImageGet(breed string) (*response.DogImageReponse, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dogResponse response.DogImageReponse
	json.Unmarshal(body, &dogResponse)

	return &dogResponse, nil
}
