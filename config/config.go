package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/caarlos0/env"
)

type Settings struct {
	Port        string `env:"PORT,required"`
	Environment string `env:"GO_ENV,required"`
	LogType     string `env:"LOG_TYPE,required"`
	LogLevel    string `env:"LOG_LEVEL,required"`
	HostUrl     string `env:"HOST_URL,required"`
	RedisUrl    string `env:"REDIS_URL,required"`
	NewRelicKey string `env:"NEW_RELIC_LICENSE_KEY,required"`
}

var config Settings

func Get() *Settings {
	return &config
}

func IsDevelopment() bool {
	return config.Environment == "development"
}

func Load() {
	LoadLogger()

	log.Info("Loading configs.")

	err := env.Parse(&config)

	if err != nil {
		log.Panicf("Fatal error config file: %s", err)
	}

	log.Info("Configs loaded successfully.")
}
