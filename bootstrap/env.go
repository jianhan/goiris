package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"

	"github.com/caarlos0/env"
)

// Env stores env variables.
type Env struct {
	AppName        string `env:"APP_NAME" envDefault:"Go App"`
	AppOwner       string `env:"APP_OWNER" envDefault:"Github Open Source"`
	AppOwnerEmail  string `env:"APP_OWNER_EMAIL" envDefault:"test@mail.com"`
	CookieHashKey  string `env:"COOKIE_HASH_KEY" envDefault:"BdHxkHTQMSIcylvWf2pkTfLPzzrZow1n9DDBfZpH"`
	CookieBlockKey string `env:"COOKIE_BLOCK_KEY" envDefault:"BTq83MeLSM4idEEGJZor6YEVbsjdLUp8iDoG6HkJ"`
	Environment    string `env:"ENVIRONMENT" envDefault:"development"`
	Host           string `env:"HOST" envDefault:"127.0.0.1"`
	Port           int    `env:"PORT" envDefault:"8888"`
	WriteTimeout   int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	ReadTimeout    int    `env:"READ_TIMEOUT" envDefault:"15"`
	IdleTimeout    int    `env:"IDLE_TIMEOUT" envDefault:"15"`
}

// Address returns derived variable for address.
func (e *Env) Address() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

// EnvConfigs load all configs.
func EnvConfigs() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("unable to load .env file", err)
		return nil, err
	}

	envConfigs := Env{}
	if err := env.Parse(&envConfigs); err != nil {
		log.Fatal(fmt.Errorf("unable to parse configs \n %+v", err))
		return nil, err
	}

	return &envConfigs, nil
}
