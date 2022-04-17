package services

import (
	"github.com/Vause/pet-information/internal/utils"
)

type Config struct {
	ConfigurationObject *utils.Configuration
}

func (c *Config) GetBreedInformation() []byte {
	return utils.GetResponseBody(c.ConfigurationObject.Url, c.ConfigurationObject.ApiKey)
}
