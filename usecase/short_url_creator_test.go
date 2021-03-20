package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/shortener/entity"
	mock_usecase "github.com/indrasaputra/shortener/test/mock/usecase"
	"github.com/indrasaputra/shortener/usecase"
)

const (
	defaultExpiryTime = 7 * 24 * time.Hour
)

type ShortURLCreatorExecutor struct {
	usecase   *usecase.ShortURLCreator
	generator *mock_usecase.MockURLGeneratorV2
	repo      *mock_usecase.MockURLRepositoryV2
}

func TestNewShortURLCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of ShortURLCreator", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func TestShortURLCreator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url can't be processed", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		urls := []string{"", "   ", "        ", "     "}

		for _, url := range urls {
			res, err := exec.usecase.Create(context.Background(), url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrEmptyURL, err)
			assert.Nil(t, res)
		}
	})

	t.Run("short url is not unique", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		original := "http://orignal-url-1.url"
		short := "http://short.url"

		exec.generator.SetReturnValues(short)
		exec.repo.SetReturnValues(entity.ErrInternalServer)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("successfully create a short url", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		original := "http://orignal-url-1.url"
		short := "http://short.url"

		exec.generator.SetReturnValues(short)
		exec.repo.SetReturnValues(nil)

		res, err := exec.usecase.Create(context.Background(), original)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func createURLEntity(original, short string) *entity.URL {
	return &entity.URL{
		ShortURL:    short,
		OriginalURL: original,
		ExpiredAt:   time.Now().Add(defaultExpiryTime),
	}
}

func createShortURLCreatorExecutor(ctrl *gomock.Controller) *ShortURLCreatorExecutor {
	g := mock_usecase.NewMockURLGeneratorV2()
	r := mock_usecase.NewMockURLRepositoryV2()
	u := usecase.NewShortURLCreator(g, r)

	return &ShortURLCreatorExecutor{
		usecase:   u,
		generator: g,
		repo:      r,
	}
}
