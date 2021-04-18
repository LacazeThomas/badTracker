package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type AppConfig struct {
	Username    string `env:"Username"`
	Password    string `env:"Password"`
	Host        string `env:"Host"`
	Port        string `env:"Port"`
	Database    string `env:"Database"`
	Environment string `default:"dev" env:"ENVIRONMENT"`
}

var appConfig AppConfig

func (a AppConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", a.Username, a.Password, a.Host, a.Port, a.Database)
}

//IsDev return true if application is on dev stack
func (a AppConfig) IsDev() bool {
	return IsDev(a.Environment)
}

//IsDev return true if application is on dev stack
func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

//InitAppConfig struct with env variables
func initAppConfig() {
	env.Parse(&appConfig)
}

//GetAppConfig return initialize config structure with variable env
func GetAppConfig() AppConfig {
	if (appConfig == AppConfig{}) {
		initAppConfig()
	}
	return appConfig
}

func SetAppConfig(cfg AppConfig) {
	appConfig = cfg
}
