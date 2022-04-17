package petInformation

import (
	"encoding/json"
	"fmt"

	"github.com/Vause/pet-information/internal/petInformation/shared"
	"github.com/Vause/pet-information/internal/services"
	"github.com/Vause/pet-information/internal/utils"
)

type Dog struct {
	BredFor          string        `json:"bred_for"`
	BreedGroup       string        `json:"breed_group"`
	Height           shared.Height `json:"height"`
	Image            shared.Image  `json:"image"`
	ID               int64         `json:"id"`
	LifeSpan         string        `json:"life_span"`
	Name             string        `json:"name"`
	Origin           string        `json:"origin"`
	ReferenceImageID string        `json:"reference_image_id"`
	Temperament      string        `json:"temperament"`
	Weight           shared.Weight `json:"weight"`
}

func GetDogBreeds(configuration *utils.Configuration) {
	configObj := services.Config{ConfigurationObject: configuration}
	body := configObj.GetBreedInformation()
	var breedResponse []Dog
	json.Unmarshal(body, &breedResponse)

	fmt.Println(breedResponse)
}
