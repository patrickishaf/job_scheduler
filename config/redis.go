package config

import "fmt"

type RedisConfig struct {
	DB       int    `envconfig:"REDIS_DB" default:"0"`
	Host     string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	Port     int    `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

func (this *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
}
