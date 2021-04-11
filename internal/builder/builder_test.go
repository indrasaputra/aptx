package builder_test

import (
	"database/sql"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
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
	t.Run("successfully build URLShortener handler", func(t *testing.T) {
		rds := &redis.Client{}
		db := &sql.DB{}
		hdr := builder.BuildGRPCURLShortener(db, rds, "http://short-url.com")
		assert.NotNil(t, hdr)
	})
}

func TestBuildGRPCHealthChecker(t *testing.T) {
	t.Run("successfully build HealthChecker handler", func(t *testing.T) {
		rds := &redis.Client{}
		db := &sql.DB{}
		hdr := builder.BuildGRPCHealthChecker(db, rds)
		assert.NotNil(t, hdr)
	})
}

func TestBuildGRPCServer(t *testing.T) {
	t.Run("successfully build gRPC server", func(t *testing.T) {
		shortener := &handler.URLShortener{}
		health := &handler.HealthChecker{}

		srv, err := builder.BuildGRPCServer("8080", shortener, health)

		assert.Nil(t, err)
		assert.NotNil(t, srv)
	})
}

func TestBuildRestServer(t *testing.T) {
	t.Run("fail build REST server due to transport security setting", func(t *testing.T) {
		srv, err := builder.BuildRestServer("8081", "8080")

		assert.NotNil(t, err)
		assert.Nil(t, srv)
	})

	t.Run("successfully build REST server", func(t *testing.T) {
		srv, err := builder.BuildRestServer("8081", "8080", grpc.WithInsecure())

		assert.Nil(t, err)
		assert.NotNil(t, srv)
	})
}
