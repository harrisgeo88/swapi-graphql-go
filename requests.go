package swapiGraphQLGo

import (
	"io/ioutil"
	"net/http"
)

// URL doc
type URL struct {
	Link string `json:"link"`
}

// Homeworld doc
type Homeworld struct {
	Homeworld string `json:"homeworld"`
}

// SendRequest doc
func SendRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// HandleRequest doc
func HandleRequest(url string) ([]byte, error) {
	response, err := SendRequest(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
