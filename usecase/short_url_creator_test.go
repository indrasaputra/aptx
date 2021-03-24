package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
	mock_usecase "github.com/indrasaputra/url-shortener/test/mock/usecase"
	"github.com/indrasaputra/url-shortener/usecase"
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

	t.Run("generator always returns error", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		original := "http://orignal-1.url"
		short := "http://short-1.url"

		exec.generator.SetReturnValues(short, entity.ErrInternalServer)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Nil(t, res)
	})

	t.Run("short url is not unique", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		original := "http://orignal-2.url"
		short := "http://short-2.url"

		exec.generator.SetReturnValues(short, nil)
		exec.repo.SetReturnValues(entity.ErrInternalServer)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("successfully create a short url", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()
		original := "http://orignal-3.url"
		short := "http://short-3.url"

		exec.generator.SetReturnValues(short, nil)
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
