package config

import (
	"context"
	"log"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type (
	AppConfig struct {
		Name        string `env:"APP_NAME"`
		Environment string `env:"APP_ENV"`
		AppPort     int    `env:"APP_PORT"`
		OauthPort   int    `env:"OAUTH_PORT"`
		DbConfig    *DbConfig
		OauthConfig *OauthConfig
	}

	OauthConfig struct {
		ClientID     string `env:"OAUTH_CLIENT_ID"`
		ClientSecret string `env:"OAUTH_SECRET_ID"`
		AuthorizeURL string `env:"OAUTH_AUTHORIZE_URL"`
		TokenURL     string `env:"OAUTH_TOKEN_URL"`
		RedirectURL  string `env:"OAUTH_REDIRECT_URL"`
	}

	DbConfig struct {
		DbHost              string `env:"DB_HOST"`
		DbPort              int    `env:"DB_PORT"`
		DbName              string `env:"DB_NAME"`
		DbUsername          string `env:"DB_USERNAME"`
		DbPassword          string `env:"DB_PASSWORD"`
		DbMaxOpenConnection int    `env:"DB_MAX_OPEN_CONNECTION"`
		DbMaxIdleConnection int    `env:"DB_MAX_IDLE_CONNECTION"`
	}
)

var Env AppConfig

func InitConfiguration() {
	ctx := context.Background()
	if err := envconfig.Process(ctx, &Env); err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(2)
	}
}
