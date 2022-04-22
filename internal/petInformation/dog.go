package petInformation

import (
	"encoding/json"

	"github.com/Vause/pet-information/internal/petInformation/shared"
	"github.com/Vause/pet-information/internal/services"
	"github.com/Vause/pet-information/internal/utils"
)

// Represents the dog object returned from the dog API
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

func GetDogs() []Dog {
	config := utils.Configurations["dog"]
	body := services.GetPetInformation(config)
	var petResponse []Dog
	json.Unmarshal(body, &petResponse)

	return petResponse
}

func (d *Dog) GetName() string {
	return d.Name
}
