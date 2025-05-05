package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost                 string `env:"DB_HOST"`
	DBPort                 string `env:"DB_PORT"`
	DBUser                 string `env:"DB_USER"`
	DBPass                 string `env:"DB_PASS"`
	DBName                 string `env:"DB_NAME"`
	ContextTimeout 			   int    `env:"CONTEXT_TIMEOUT"`
	ServerAddress          string `env:"SERVER_ADDRESS"`
	AccessTokenExpiryHour  int    `env:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `env:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `env:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigName(".emv")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}