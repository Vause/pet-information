package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	ApiKey string
	Url    string
}

func GetConfig() *Configuration {
	configFile := "app_config.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}

	return &Configuration{
		ApiKey: viper.GetString("app.api.key"),
		Url:    viper.GetString("app.api.url"),
	}
}
