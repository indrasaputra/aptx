package builder_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
)

func TestBuildRedisClient(t *testing.T) {
	t.Run("fail build redis client", func(t *testing.T) {
		server, _ := miniredis.Run()

		cfg := config.Redis{
			Address: server.Addr(),
		}

		server.Close()
		client, err := builder.BuildRedisClient(cfg)

		assert.NotNil(t, err)
		assert.Nil(t, client)
	})

	t.Run("success build redis client", func(t *testing.T) {
		server, _ := miniredis.Run()
		defer server.Close()

		cfg := config.Redis{
			Address: server.Addr(),
		}

		client, err := builder.BuildRedisClient(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, client)
	})
}

func TestBuildPostgresClient(t *testing.T) {
	t.Run("success build db client", func(t *testing.T) {
		cfg := config.Postgres{
			Host:         "localhost",
			Port:         "5432",
			DBName:       "url_shortener",
			User:         "user",
			Password:     "password",
			MaxOpenConns: 10,
			MaxIdleConns: 10,
		}

		client, err := builder.BuildPostgresClient(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, client)
	})
}

func TestBuildGRPCURLShortener(t *testing.T) {
	t.Run("successfully create URLShortener handler", func(t *testing.T) {
		hdr := builder.BuildGRPCURLShortener("http://short-url.com")
		assert.NotNil(t, hdr)
	})
}

func TestBuildGRPCHealthChecker(t *testing.T) {
	t.Run("successfully create HealthChecker handler", func(t *testing.T) {
		hdr := builder.BuildGRPCHealthChecker()
		assert.NotNil(t, hdr)
	})
}
