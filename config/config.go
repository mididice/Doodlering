package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var DBEnv Database

type Database struct {
	Type     string
	User     string
	Host     string
	Name     string
	Password string
}

func SetupEnv() error {
	rootDir, _ := filepath.Abs("./")
	viper.SetConfigFile(rootDir + "/config/config.json")

	var err error
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error read config file, %s", err)
	}

	database := viper.Sub("database")
	err = database.Unmarshal(&DBEnv)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return err
}
