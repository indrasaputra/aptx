package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
	mock_usecase "github.com/indrasaputra/url-shortener/test/mock/usecase"
	"github.com/indrasaputra/url-shortener/usecase"
)

var (
	testContext            = context.Background()
	testURLCode            = "AbCdE12"
	testURLShort           = "http://localhost/" + testURLCode
	testURLOriginal        = "http://very-long-url.url"
	testErrInternalMessage = "unexpected command"
)

type ShortURLCreatorExecutor struct {
	usecase   *usecase.ShortURLCreator
	generator *mock_usecase.MockURLGeneratorV2
	repo      *mock_usecase.MockCreateShortURLRepositoryV2
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
			res, err := exec.usecase.Create(testContext, url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrEmptyURL(), err)
			assert.Nil(t, res)
		}
	})

	t.Run("generator always returns error", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()

		exec.generator.SetReturnValues("", "", entity.ErrInternal(testErrInternalMessage))

		res, err := exec.usecase.Create(testContext, testURLOriginal)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Nil(t, res)
	})

	t.Run("short url is not unique", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()

		exec.generator.SetReturnValues(testURLCode, testURLShort, nil)
		exec.repo.SetReturnValues(entity.ErrInternal(testErrInternalMessage))

		res, err := exec.usecase.Create(testContext, testURLOriginal)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("successfully create a short url", func(t *testing.T) {
		exec := createShortURLCreatorExecutor()

		exec.generator.SetReturnValues(testURLCode, testURLShort, nil)
		exec.repo.SetReturnValues(nil)

		res, err := exec.usecase.Create(testContext, testURLOriginal)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func createShortURLCreatorExecutor() *ShortURLCreatorExecutor {
	g := mock_usecase.NewMockURLGeneratorV2()
	r := mock_usecase.NewMockCreateShortURLRepositoryV2()
	u := usecase.NewShortURLCreator(g, r)

	return &ShortURLCreatorExecutor{
		usecase:   u,
		generator: g,
		repo:      r,
	}
}
