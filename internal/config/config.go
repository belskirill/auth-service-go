package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env   string `env:"ENV,required"`
	GRPC  GRPCConfig
	Token TokenConfig
}

type TokenConfig struct {
	TokenTTL       time.Duration `env:"TOKEN_TTL,required"`
	TokenSecretKey string        `env:"TOKEN_SECRET_KEY,required"`
}

type GRPCConfig struct {
	Port    int           `env:"PORT_GRPC,required"`
	Timeout time.Duration `env:"GRPC_TIMEOUT,required"`
}

func LoadConfig() Config {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	log.Println("loaded config")
	return cfg
}
