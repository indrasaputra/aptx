package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/shortener/entity"
	mock_usecase "github.com/indrasaputra/shortener/test/mock/usecase"
	"github.com/indrasaputra/shortener/usecase"
)

type ShortURLCreatorExecutor struct {
	usecase   *usecase.ShortURLCreator
	generator *mock_usecase.MockURLGeneratorV2
	repo      *mock_usecase.MockURLRepositoryV2
}

func TestNewShortURLCreator(t *testing.T) {
	t.Run("successfully create an instance of ShortURLCreator", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		assert.NotNil(t, exec.usecase)
	})
}

func TestShortURLCreator_Create(t *testing.T) {
	t.Run("empty url can't be processed", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		urls := []string{"", "   ", "        ", "     "}

		for _, url := range urls {
			res, err := exec.usecase.Create(context.Background(), url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrEmptyURL, err)
			assert.Nil(t, res)
		}
	})

	t.Run("short url is not unique", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		original := "http://orignal-url-1.url"
		short := "http://short.url"

		exec.generator.SetReturnValues(short)
		exec.repo.SetReturnValues(entity.ErrInternalServer)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("successfully create a short url", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		original := "http://orignal-url-1.url"
		short := "http://short.url"

		exec.generator.SetReturnValues(short)
		exec.repo.SetReturnValues(nil)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func createShortURLCreatorExecutor() *ShortURLCreatorExecutor {
	g := mock_usecase.NewMockURLGeneratorV2()
	r := mock_usecase.NewMockURLRepositoryV2()
	u := usecase.NewShortURLCreator(g, r)

	return &ShortURLCreatorExecutor{
		usecase:   u,
		generator: g,
		repo:      r,
	}
}
