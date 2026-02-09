package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string `env:"ENV,required"`
	GRPC     GRPCConfig
	Token    TokenConfig
	DataBase DataBaseConfig
}

type TokenConfig struct {
	TokenTTL       time.Duration `env:"TOKEN_TTL,required"`
	TokenSecretKey string        `env:"TOKEN_SECRET_KEY,required"`
}

type DataBaseConfig struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	SSLMode  string `env:"DB_SSL_MODE"`
}

func (c DataBaseConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode,
	)
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
