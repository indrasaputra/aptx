package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

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
