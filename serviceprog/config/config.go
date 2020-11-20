package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("development")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
