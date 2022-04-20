package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Key string `mapstructure:"key"`
	Url string `mapstructure:"url"`
}

var Configurations map[string]*Config

func SetUpConfig() {
	configFile := "app_config.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
	viper.UnmarshalKey("app", &Configurations)
}
