package petInformation

import (
	"encoding/json"

	"github.com/Vause/pet-information/internal/petInformation/shared"
	"github.com/Vause/pet-information/internal/services"
	"github.com/Vause/pet-information/internal/utils"
)

type Cat struct {
	Adaptability     int64         `json:"adaptability"`
	AffectionLevel   int64         `json:"affection_level"`
	Bidability       int64         `json:"bidability"`
	CatFriendly      int64         `json:"cat_friendly"`
	CfaURL           string        `json:"cfa_url"`
	ChildFriendly    int64         `json:"child_friendly"`
	CountryCode      string        `json:"country_code"`
	CountryCodes     string        `json:"country_codes"`
	Description      string        `json:"description"`
	DogFriendly      int64         `json:"dog_friendly"`
	EnergyLevel      int64         `json:"energy_level"`
	Experimental     int64         `json:"experimental"`
	Grooming         int64         `json:"grooming"`
	Hairless         int64         `json:"hairless"`
	HealthIssues     int64         `json:"health_issues"`
	Height           shared.Height `json:"height"`
	Hypoallergenic   int64         `json:"hypoallergenic"`
	ID               string        `json:"id"`
	Image            shared.Image  `json:"image"`
	Indoor           int64         `json:"indoor"`
	Intelligence     int64         `json:"intelligence"`
	Lap              int64         `json:"lap"`
	LifeSpan         string        `json:"life_span"`
	Name             string        `json:"name"`
	Natural          int64         `json:"natural"`
	Origin           string        `json:"origin"`
	Rare             int64         `json:"rare"`
	ReferenceImageID string        `json:"reference_image_id"`
	Rex              int64         `json:"rex"`
	SheddingLevel    int64         `json:"shedding_level"`
	ShortLegs        int64         `json:"short_legs"`
	SocialNeeds      int64         `json:"social_needs"`
	StrangerFriendly int64         `json:"stranger_friendly"`
	SuppressedTail   int64         `json:"suppressed_tail"`
	Temperament      string        `json:"temperament"`
	VcahospitalsURL  string        `json:"vcahospitals_url"`
	VetstreetURL     string        `json:"vetstreet_url"`
	Vocalisation     int64         `json:"vocalisation"`
	Weight           shared.Weight `json:"weight"`
	WikipediaURL     string        `json:"wikipedia_url"`
}

func GetCats() []Cat {
	config := utils.Configurations["cat"]
	body := services.GetPetInformation(config)
	var breedResponse []Cat
	json.Unmarshal(body, &breedResponse)

	return breedResponse
}

func (c Cat) GetName() string {
	return c.Name
}
