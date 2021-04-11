package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Redis holds configuration for the redis.
type Redis struct {
	Address  string `env:"REDIS_ADDRESS,required"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}

// Postgres holds configuration for PostgreSQL.
type Postgres struct {
	Host            string `env:"POSTGRES_HOST,required"`
	Port            string `env:"POSTGRES_PORT,required"`
	DBName          string `env:"POSTGRES_DBNAME,required"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	MaxOpenConns    int    `env:"POSTGRES_MAX_OPEN_CONNS,required"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN_LIFETIME,required"`
}

// Config holds configuration for the project.
type Config struct {
	PortHTTP string `env:"PORT_HTTP,default=8081"`
	PortGRPC string `env:"PORT_GRPC,default=8080"`
	Domain   string `env:"DOMAIN,required"`
	Postgres Postgres
	Redis    Redis
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
