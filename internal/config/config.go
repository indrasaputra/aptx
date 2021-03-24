package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Port   string `env:"PORT,default=8080"`
	Domain string `env:"DOMAIN,required"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
