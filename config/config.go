package config

import (
	"github.com/tkanos/gonfig"
)

type DBConfiguration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() DBConfiguration {
	configuration := DBConfiguration{}
	fileName := "config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
