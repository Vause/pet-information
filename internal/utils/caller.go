package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResponseBody(url, key string) []byte {
	timeoutClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, _ := http.NewRequest("GET", url, nil)

	if key != "" {
		req.Header.Add("X-API-KEY", key)
	}

	res, err := timeoutClient.Do(req)

	if err != nil {
		errors.New("Bad Response")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		errors.New("Unable To Read Body")
	}

	return body
}
