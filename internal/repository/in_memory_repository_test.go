package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/repository"
)

var (
	testCodes = []string{"ABCdef12", "xyzJKL34", "asdQWE56"}
)

func TestNewInMemoryURLRepository(t *testing.T) {
	t.Run("successfully create an instance of InMemoryURLRepository", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		assert.NotNil(t, repo)
	})
}

func TestInMemoryURLRepository_Save(t *testing.T) {
	t.Run("can't save duplicated short URL", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		for i, code := range testCodes {
			url := createURL(fmt.Sprintf("http://original-%d.url", i), code)
			repo.Save(context.Background(), url)
		}

		for i, code := range testCodes {
			url := createURL(fmt.Sprintf("http://original-%d-%d.url", i, i), code)
			err := repo.Save(context.Background(), url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrAlreadyExists(), err)
		}
	})

	t.Run("success save short url", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		for i, code := range testCodes {
			url := createURL(fmt.Sprintf("http://original-%d.url", i), code)
			err := repo.Save(context.Background(), url)
			assert.Nil(t, err)
		}
	})
}

func TestInMemoryURLRepository_GetAll(t *testing.T) {
	t.Run("empty data returns empty list and nil error", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()

		urls, err := repo.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Empty(t, urls)
	})

	t.Run("repository returns as many as data previously stored", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		fillRepository(repo, 10)

		urls, err := repo.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 10, len(urls))
	})
}

func TestInMemoryURLRepository_GetByCode(t *testing.T) {
	t.Run("wanted URL doesn't exist", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		fillRepository(repo, 10)

		url, err := repo.GetByCode(context.Background(), "http://not-found-short.url")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, url)
	})

	t.Run("successfully get single url", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		fillRepository(repo, 10)

		url, err := repo.GetByCode(context.Background(), "random-1")

		assert.Nil(t, err)
		assert.NotNil(t, url)
		assert.Equal(t, "http://original-random-1.url", url.OriginalURL)
		assert.Equal(t, "http://short.url/random-1", url.ShortURL)
	})
}

func fillRepository(repo *repository.InMemoryURLRepository, numberOfURL int) {
	for i := 0; i < numberOfURL; i++ {
		repo.Save(context.Background(), createURL(fmt.Sprintf("http://original-random-%d.url", i), fmt.Sprintf("random-%d", i)))
	}
}

func createURL(original, code string) *entity.URL {
	return &entity.URL{
		Code:        code,
		ShortURL:    "http://short.url/" + code,
		OriginalURL: original,
		ExpiredAt:   time.Now().Add(1 * time.Hour),
	}
}
