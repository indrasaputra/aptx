package handler_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
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

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL, err)
		assert.Nil(t, resp)
	})

	t.Run("creator usecase returns error", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		req := &shortenerv1.CreateShortURLRequest{OriginalUrl: "http://original-1.url"}
		exec.usecase.EXPECT().Create(context.Background(), req.GetOriginalUrl()).Return(nil, entity.ErrInternalServer)

		resp, err := exec.handler.CreateShortURL(context.Background(), req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Nil(t, resp)
	})

	t.Run("successfully create a shorturl", func(t *testing.T) {
		exec := createShortURLCreatorExecutor(ctrl)
		req := &shortenerv1.CreateShortURLRequest{OriginalUrl: "http://original-1.url"}
		now := time.Now()
		url := &entity.URL{ShortURL: "http://short-1.url", ExpiredAt: now}
		exec.usecase.EXPECT().Create(context.Background(), req.GetOriginalUrl()).Return(url, nil)

		resp, err := exec.handler.CreateShortURL(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "http://short-1.url", resp.GetShortUrl())
		assert.Equal(t, timestamppb.New(now), resp.GetExpiredAt())
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
