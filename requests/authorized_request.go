package requests

import (
	"io"
	"log"
	"net/http"
)

func AuthorizedGetRequest(url string, bearerToken string) (*http.Response, error) {
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearerToken)

	// Send req using http Client
	client := &http.Client{}
	return client.Do(req)
}

func AuthorizedBodyRequest(url string, method string, bearerToken string, body io.Reader) (*http.Response, error) {
	// Create a new request using http
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearerToken)

	// Send req using http Client
	client := &http.Client{}
	return client.Do(req)
}
