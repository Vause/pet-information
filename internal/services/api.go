package services

import (
	"github.com/Vause/pet-information/internal/utils"
)

func GetBreedInformation(config *utils.Config) []byte {
	return utils.GetResponseBody(config.Url, config.Key)
}
