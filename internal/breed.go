package internal

import (
	"encoding/json"
	"fmt"

	"github.com/Vause/pet-information/internal/services"
)

type Breed struct {
	BredFor          string `json:"bred_for"`
	BreedGroup       string `json:"breed_group"`
	Height           Height `json:"height"`
	Image            Image  `json:"image"`
	ID               int64  `json:"id"`
	LifeSpan         string `json:"life_span"`
	Name             string `json:"name"`
	Origin           string `json:"origin"`
	ReferenceImageID string `json:"reference_image_id"`
	Temperament      string `json:"temperament"`
	Weight           Weight `json:"weight"`
}

type Image struct {
	Height int64  `json:"height"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Height struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

type Weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

func GetBreeds() {
	body := services.GetResponseBody()
	var breedResponse []Breed
	json.Unmarshal(body, &breedResponse)

	fmt.Println(breedResponse)
}
