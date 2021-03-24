package handler_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	mock_usecase "github.com/indrasaputra/url-shortener/test/mock/usecase"
)

type ShortURLCreatorExecutor struct {
	handler *handler.ShortURLCreator
	usecase *mock_usecase.MockCreateShortURL
}

func TestNewShortURLCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of ShortURLCreator", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestShortURLCreator_CreateShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)

		resp, err := exec.handler.CreateShortURL(context.Background(), nil)

		assert.NotNil(t, exec.handler)
		assert.NotNil(t, entity.ErrEmptyURL, err)
		assert.Nil(t, resp)
	})
}

func createShortURLCreatorExecutor(ctrl *gomock.Controller) *ShortURLCreatorExecutor {
	c := mock_usecase.NewMockCreateShortURL(ctrl)
	h := handler.NewShortURLCreator(c)
	return &ShortURLCreatorExecutor{
		handler: h,
		usecase: c,
	}
}
