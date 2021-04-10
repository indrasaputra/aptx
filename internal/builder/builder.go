package builder

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq" // for posgres

	"github.com/indrasaputra/url-shortener/internal/config"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	"github.com/indrasaputra/url-shortener/internal/repository"
	"github.com/indrasaputra/url-shortener/internal/repository/cache"
	"github.com/indrasaputra/url-shortener/internal/repository/database"
	"github.com/indrasaputra/url-shortener/internal/tool"
	"github.com/indrasaputra/url-shortener/usecase"
)

const (
	dbDriver = "postgres"
)

// BuildRedisClient builds a redis client.
func BuildRedisClient(cfg config.Redis) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client := redis.NewClient(opt)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// BuildPostgresClient builds a postgres client.
func BuildPostgresClient(cfg config.Postgres) (*sql.DB, error) {
	sqlCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open(dbDriver, sqlCfg)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}

// BuildGRPCURLShortener builds URLShortener handler together with all of its dependencies.
func BuildGRPCURLShortener(db *sql.DB, rds redis.Cmdable, domain string) *handler.URLShortener {
	redis := cache.NewURLRedis(rds)
	postgres := database.NewURLPostgres(db)

	repoInserter := repository.NewURLInserter(postgres, redis)
	repoGetter := repository.NewURLGetter(postgres, redis)

	gen := tool.NewShortURLGenerator(domain)

	creator := usecase.NewShortURLCreator(gen, repoInserter)
	getter := usecase.NewURLGetter(repoGetter)

	return handler.NewURLShortener(creator, getter)
}

// BuildGRPCHealthChecker builds HealthChecker handler together with all of its dependencies.
func BuildGRPCHealthChecker(db *sql.DB, rds redis.Cmdable) *handler.HealthChecker {
	redis := cache.NewURLRedis(rds)
	postgres := database.NewURLPostgres(db)

	repo := repository.NewHealthChecker(postgres, redis)

	checker := usecase.NewHealthChecker(repo)
	return handler.NewHealthChecker(checker)
}
