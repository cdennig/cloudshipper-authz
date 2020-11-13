package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds data for application configuration
type Config struct {
	DbURL  string `required:"true"`
	DbKey  string `required:"true"`
	DbName string `required:"true"`
	Port   string `required:"true"`
	Env    string
}

// Load returns config from yaml and environment variables.
func Load() (*Config, error) {
	log.Print("Loading config from ENV.")

	// Load from .env file, if not 'prod'
	if os.Getenv("CS_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	cfg := Config{}

	if err := envconfig.Process("cs", &cfg); err != nil {
		log.Printf("Error loading config: %s", err)
		return nil, err
	}

	return &cfg, nil
}
