package app

import (
	"github.com/netflix/go-env"
)

var ENV *Environment

type Environment struct {
	DBName             string `env:"DB_NAME"`
	DBHost             string `env:"DB_HOST"`
	DBPort             int    `env:"DB_PORT"`
	DBUser             string `env:"DB_USER"`
	DBPassword         string `env:"DB_PASSWORD"`
	MaxPoolConnections int    `env:"MAX_POOL_CONNECTIONS,default=10"`
	AppPort            int    `env:"APP_PORT,default=8080"`
}

func LoadVars() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}
	ENV = &environment
}
