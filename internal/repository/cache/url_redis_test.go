package cache_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/repository/cache"
)

var (
	testContext      = context.Background()
	testURLCode      = "AbCdE12"
	testURLShort     = "http://localhost/" + testURLCode
	testURLOriginal  = "http://very-long-url.url"
	testURLExpiredAt = time.Now().Add(1 * time.Minute)
	testURLCreatedAt = time.Now()
	testURL          = &entity.URL{
		Code:        testURLCode,
		ShortURL:    testURLShort,
		OriginalURL: testURLOriginal,
		ExpiredAt:   testURLExpiredAt,
		CreatedAt:   testURLCreatedAt,
	}
	testEmptyMapResult = make(map[string]string)
	testValidMapResult = map[string]string{
		"code":         testURLCode,
		"short_url":    testURLShort,
		"original_url": testURLOriginal,
		"expired_at":   testURLExpiredAt.Format(time.RFC3339),
		"created_at":   testURLCreatedAt.Format(time.RFC3339),
	}
	testHSetInput = []string{
		"code",
		testURLCode,
		"short_url",
		testURLShort,
		"original_url",
		testURLOriginal,
		"expired_at",
		testURLExpiredAt.Format(time.RFC3339),
		"created_at",
		testURLCreatedAt.Format(time.RFC3339),
	}
	testRedisDownMessage = "redis down"
)

type URLRedisExecutor struct {
	redis *cache.URLRedis
	mock  redismock.ClientMock
}

func TestNewURLRedis(t *testing.T) {
	t.Run("successfully create an instance of Redis", func(t *testing.T) {
		exec := createURLRedisExecutor()
		assert.NotNil(t, exec.redis)
	})
}

func TestURLRedis_Save(t *testing.T) {
	t.Run("not all attributes are saved", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHSet(testURLCode, testHSetInput).SetVal(2)
		exec.mock.ExpectExpireAt(testURLCode, testURLExpiredAt).SetVal(true)

		err := exec.redis.Save(testContext, testURL)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "only success to save 2 out of 5 attributes")
	})

	t.Run("redis is down", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHSet(testURLCode, testHSetInput).SetVal(5)
		exec.mock.ExpectExpireAt(testURLCode, testURLExpiredAt).SetErr(errors.New(testRedisDownMessage))

		err := exec.redis.Save(testContext, testURL)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), testRedisDownMessage)
	})

	t.Run("success save url in redis hash", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHSet(testURLCode, testHSetInput).SetVal(5)
		exec.mock.ExpectExpireAt(testURLCode, testURLExpiredAt).SetVal(true)

		err := exec.redis.Save(testContext, testURL)

		assert.Nil(t, err)
	})
}

func TestURLRedis_Get(t *testing.T) {
	t.Run("redis hgetall returns error", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHGetAll(testURLCode).SetErr(errors.New(testRedisDownMessage))

		url, err := exec.redis.Get(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, url)
	})

	t.Run("redis hgetall returns empty hash", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHGetAll(testURLCode).SetVal(testEmptyMapResult)

		url, err := exec.redis.Get(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, url)
	})

	t.Run("url's expired_at is invalid", func(t *testing.T) {
		exec := createURLRedisExecutor()
		hash := make(map[string]string)
		hash["expired_at"] = ""
		exec.mock.ExpectHGetAll(testURLCode).SetVal(hash)

		url, err := exec.redis.Get(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, url)
	})

	t.Run("url's created_at is invalid", func(t *testing.T) {
		exec := createURLRedisExecutor()
		hash := make(map[string]string)
		hash["expired_at"] = "2021-04-04T12:05:38.728727+07:00"
		hash["created_at"] = ""
		exec.mock.ExpectHGetAll(testURLCode).SetVal(hash)

		url, err := exec.redis.Get(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, url)
	})

	t.Run("success get url from redis hash", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectHGetAll(testURLCode).SetVal(testValidMapResult)

		url, err := exec.redis.Get(testContext, testURLCode)
		fmt.Println(err)

		assert.Nil(t, err)
		assert.NotNil(t, url)
	})
}

func TestURLRedis_IsAlive(t *testing.T) {
	t.Run("IsAlive returns error", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectPing().SetErr(errors.New("redis: nil"))

		res := exec.redis.IsAlive(context.Background())

		assert.False(t, res)
	})

	t.Run("IsAlive doesn't return PONG", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectPing().SetVal("PINGPONG")

		res := exec.redis.IsAlive(context.Background())

		assert.False(t, res)
	})

	t.Run("redis is alive", func(t *testing.T) {
		exec := createURLRedisExecutor()
		exec.mock.ExpectPing().SetVal("PONG")

		res := exec.redis.IsAlive(context.Background())

		assert.True(t, res)
	})
}

func createURLRedisExecutor() *URLRedisExecutor {
	client, mock := redismock.NewClientMock()
	rds := cache.NewURLRedis(client)
	// mock.Exc
	return &URLRedisExecutor{
		redis: rds,
		mock:  mock,
	}
}
