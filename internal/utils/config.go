package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigurationSections struct {
	Config map[string]ConfigurationSection
}

type ConfigurationSection struct {
	Key string
	Url string
}

type App struct {
	Cat ConfigurationSection
	Dog ConfigurationSection
}

type Config struct {
	App App
}

var ConfigSections *ConfigurationSections

func SetUpConfig() {
	configFile := "app_config.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}

	var cfg Config

	viper.Unmarshal(&cfg)

	var m = make(map[string]ConfigurationSection)
	m["cat"] = cfg.App.Cat
	m["dog"] = cfg.App.Dog
	ConfigSections = &ConfigurationSections{Config: m}
}
