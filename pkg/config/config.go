package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitViper() {
	workDir, _ := os.Getwd()
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Config not Found")
		}
		panic(err)
	}

	log.Print("Config Read ok")
}
