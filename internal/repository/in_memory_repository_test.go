package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/shortener/entity"
	"github.com/indrasaputra/shortener/internal/repository"
)

func TestNewInMemoryURLRepository(t *testing.T) {
	t.Run("successfully create an instance of InMemoryURLRepository", func(t *testing.T) {
		repo := repository.NewInMemoryURLRepository()
		assert.NotNil(t, repo)
	})
}

func TestInMemoryURLRepository_Save(t *testing.T) {
	t.Run("can't save duplicated short URL", func(t *testing.T) {
		shorts := []string{"http://short-1.url", "http://short-2.url", "http://short-3.url"}

		repo := repository.NewInMemoryURLRepository()
		for i, short := range shorts {
			url := createURL(fmt.Sprintf("http://original-%d.url", i), short)
			repo.Save(context.Background(), url)
		}

		for i, short := range shorts {
			url := createURL(fmt.Sprintf("http://original-%d-%d.url", i, i), short)
			err := repo.Save(context.Background(), url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrDuplicatedShortURL, err)
		}
	})

	t.Run("success save short url", func(t *testing.T) {
		shorts := []string{"http://short-1.url", "http://short-2.url", "http://short-3.url"}

		repo := repository.NewInMemoryURLRepository()
		for i, short := range shorts {
			url := createURL(fmt.Sprintf("http://original-%d.url", i), short)
			err := repo.Save(context.Background(), url)
			assert.Nil(t, err)
		}
	})
}

func createURL(original, short string) *entity.URL {
	return &entity.URL{
		ShortURL:    short,
		OriginalURL: original,
		ExpiredAt:   time.Now().Add(1 * time.Hour),
	}
}
