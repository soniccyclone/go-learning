package main

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connectionString"`
}

func LoadAppConfig() (*config, error) {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var appConfig *config
	err = viper.Unmarshal(&appConfig)
	if err == nil {
		return appConfig, nil
	}
	return nil, err
}
