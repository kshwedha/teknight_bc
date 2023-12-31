package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	log.Print("This is the environment: ", env)
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	config.AddConfigPath(".")             // used for docker
	config.AddConfigPath("../../config/") // used for unit tests
	err = config.ReadInConfig()
	if err != nil {
		log.Fatalf("error on parsing configuration file: %s", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
