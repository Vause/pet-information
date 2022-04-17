package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Vause/pet-information/internal/utils"
)

var (
	url = ""
	key = ""
)

func SetUp(configuration *utils.Configuration) {
	url = configuration.Url
	key = configuration.ApiKey
}

func GetResponseBody() []byte {
	timeoutClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-API-KEY", key)

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
