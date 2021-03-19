package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/shortener/entity"
	mock_usecase "github.com/indrasaputra/shortener/test/mock/usecase"
	"github.com/indrasaputra/shortener/usecase"
)

type ShortURLCreatorExecutor struct {
	usecase   *usecase.ShortURLCreator
	generator *mock_usecase.MockURLGenerator
	repo      *mock_usecase.MockURLRepository
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

	t.Run("empty url can't be processedd", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		urls := []string{"", "   ", "        ", "     "}

		for _, url := range urls {
			res, err := exec.usecase.Create(context.Background(), url)

			assert.NotNil(t, err)
			assert.Equal(t, entity.ErrEmptyURL, err)
			assert.Nil(t, res)
		}
	})
}

func createShortURLCreatorExecutor(ctrl *gomock.Controller) *ShortURLCreatorExecutor {
	g := mock_usecase.NewMockURLGenerator(ctrl)
	r := mock_usecase.NewMockURLRepository(ctrl)
	u := usecase.NewShortURLCreator(g, r)

	return &ShortURLCreatorExecutor{
		usecase:   u,
		generator: g,
		repo:      r,
	}
}
