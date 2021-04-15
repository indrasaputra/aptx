package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/indrasaputra/aptx/entity"
)

const (
	redisPong = "PONG"
)

var (
	attributes        = []string{"code", "short_url", "original_url", "_expired_at", "created_at"}
	numberOfAttribute = len(attributes)
)

// URLRedis is responsible to connect with Redis.
// It uses https://github.com/go-redis/redis.
type URLRedis struct {
	client redis.Cmdable
}

// NewURLRedis creates an instance URLRedis.
func NewURLRedis(client redis.Cmdable) *URLRedis {
	return &URLRedis{client: client}
}

// Save saves URL in Redis.
// It uses hash (https://redis.io/commands/hset).
func (ur *URLRedis) Save(ctx context.Context, url *entity.URL) error {
	hash := createURLHash(url)

	pipe := ur.client.Pipeline()
	res := pipe.HSet(ctx, url.Code, hash)
	pipe.ExpireAt(ctx, url.Code, url.ExpiredAt)
	_, err := pipe.Exec(ctx)

	if int(res.Val()) != numberOfAttribute {
		return entity.ErrInternal(fmt.Sprintf("only success to save %d out of %d attributes", res.Val(), numberOfAttribute))
	}
	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	return nil
}

// Get gets URL detail from Redis.
func (ur *URLRedis) Get(ctx context.Context, key string) (*entity.URL, error) {
	hash, err := ur.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	if len(hash) == 0 {
		return nil, entity.ErrNotFound()
	}
	return createURLFromHash(hash)
}

// IsAlive must returns true if Redis can connect without any problem.
// It basically calls PING command.
func (ur *URLRedis) IsAlive(ctx context.Context) bool {
	res, err := ur.client.Ping(ctx).Result()
	return err == nil && res == redisPong
}

func createURLFromHash(hash map[string]string) (*entity.URL, error) {
	url := &entity.URL{}
	var err error

	url.Code = hash["code"]
	url.ShortURL = hash["short_url"]
	url.OriginalURL = hash["original_url"]
	url.ExpiredAt, err = time.Parse(time.RFC3339, hash["expired_at"])
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	url.CreatedAt, err = time.Parse(time.RFC3339, hash["created_at"])
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	return url, nil
}

func createURLHash(url *entity.URL) []string {
	return []string{
		"code",
		url.Code,
		"short_url",
		url.ShortURL,
		"original_url",
		url.OriginalURL,
		"expired_at",
		url.ExpiredAt.Format(time.RFC3339),
		"created_at",
		url.CreatedAt.Format(time.RFC3339),
	}
}
