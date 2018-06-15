package kiwi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const apiURL = "https://api.skypicker.com"

var (
	// ErrorStatusCodeNotOK when the status code is not 200
	ErrorStatusCodeNotOK = errors.New("Status code not 200")
)

// Kiwi structure that handles the different calls to the API
type Kiwi struct {
	endpoint string
}

// New creates a new Kiwi structure with the default endpoint
func New() *Kiwi {
	return &Kiwi{endpoint: apiURL}
}

// Call executed an API call to the kiwi endpoint
func (k *Kiwi) call(path string, v interface{}) error {
	resp, err := http.Get(k.endpoint + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrorStatusCodeNotOK
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
