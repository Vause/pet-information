package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	ApiKey string
	Url    string
}

func GetConfig(section string) *Configuration {
	configFile := "app_config.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}

	apiKey := fmt.Sprintf("app.%s.key", section)
	url := fmt.Sprintf("app.%s.url", section)

	return &Configuration{
		ApiKey: viper.GetString(apiKey),
		Url:    viper.GetString(url),
	}
}
