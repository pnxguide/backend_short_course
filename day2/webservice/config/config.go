package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigFile("yaml")
	config.SetConfigName("development")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
