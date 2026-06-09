package config

import (
	"log"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

var config Config

func LoadConfig() *Config {
	godotenv.Load()
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("failed to load env %v", err)
	}
	return &config
}

func GetConfig() *Config {
	if reflect.DeepEqual(config, Config{}) {
		return LoadConfig()
	}
	return &config
}
