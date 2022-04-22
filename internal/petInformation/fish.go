package petInformation

import (
	"encoding/json"

	"github.com/Vause/pet-information/internal/services"
	"github.com/Vause/pet-information/internal/utils"
)

// Represents the fish object returned from the fish API
type Fish struct {
	AvLength    int64   `json:"avLength"`
	Biology     string  `json:"biology"`
	Description string  `json:"description"`
	Environment string  `json:"environment"`
	ID          int64   `json:"id"`
	MaxAge      int64   `json:"maxAge"`
	MaxClimate  int64   `json:"maxClimate"`
	MaxDepth    int64   `json:"maxDepth"`
	MaxLength   int64   `json:"maxLength"`
	MaxWeight   float64 `json:"maxWeight"`
	MinClimate  int64   `json:"minClimate"`
	MinDepth    int64   `json:"minDepth"`
	Name        string  `json:"name"`
	Scientific  string  `json:"scientific"`
}

func GetFishes() []Fish {
	config := utils.Configurations["fish"]
	body := services.GetPetInformation(config)
	var petResponse []Fish
	json.Unmarshal(body, &petResponse)

	return petResponse
}

func (f *Fish) GetName() string {
	return f.Name
}
