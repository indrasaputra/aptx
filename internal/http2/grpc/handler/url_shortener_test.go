package handler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
	mock_grpc "github.com/indrasaputra/url-shortener/test/mock/http2/grpc"
	mock_usecase "github.com/indrasaputra/url-shortener/test/mock/usecase"
)

var (
	globalURLs = []*entity.URL{
		{
			ShortURL:    "http://short-1.url",
			OriginalURL: "http://original-1.url",
			ExpiredAt:   time.Now().Add(1 * time.Minute),
		},
		{
			ShortURL:    "http://short-2.url",
			OriginalURL: "http://original-2.url",
			ExpiredAt:   time.Now().Add(2 * time.Minute),
		},
		{
			ShortURL:    "http://short-3.url",
			OriginalURL: "http://original-3.url",
			ExpiredAt:   time.Now().Add(2 * time.Minute),
		},
	}

	globalsResponses = createGetAllURLReponse(globalURLs)
)

type URLShortenerExecutor struct {
	handler *handler.URLShortener
	creator *mock_usecase.MockCreateShortURL
	getter  *mock_usecase.MockGetURL
}

func TestNewURLShortener(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of URLShortener", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestURLShortener_CreateShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)

		resp, err := exec.handler.CreateShortURL(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL, err)
		assert.Nil(t, resp)
	})

	t.Run("creator usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.CreateShortURLRequest{OriginalUrl: "http://original-1.url"}
		exec.creator.EXPECT().Create(context.Background(), req.GetOriginalUrl()).Return(nil, entity.ErrInternalServer)

		resp, err := exec.handler.CreateShortURL(context.Background(), req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Nil(t, resp)
	})

	t.Run("successfully create a shorturl", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.CreateShortURLRequest{OriginalUrl: "http://original-1.url"}
		now := time.Now()
		url := &entity.URL{ShortURL: "http://short-1.url", ExpiredAt: now}
		exec.creator.EXPECT().Create(context.Background(), req.GetOriginalUrl()).Return(url, nil)

		resp, err := exec.handler.CreateShortURL(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "http://short-1.url", resp.GetShortUrl())
		assert.Equal(t, timestamppb.New(now), resp.GetExpiredAt())
	})
}

func TestURLShortener_GetAllURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetAllURLRequest{}
		stream := mock_grpc.NewMockURLShortenerService_GetAllURLServer(ctrl)
		exec.getter.EXPECT().GetAll(context.Background()).Return([]*entity.URL{}, entity.ErrInternalServer)

		err := exec.handler.GetAllURL(req, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream can't send response", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetAllURLRequest{}
		resp := &shortenerv1.GetAllURLResponse{
			ShortUrl:    globalURLs[0].ShortURL,
			OriginalUrl: globalURLs[0].OriginalURL,
			ExpiredAt:   timestamppb.New(globalURLs[0].ExpiredAt),
		}
		stream := mock_grpc.NewMockURLShortenerService_GetAllURLServer(ctrl)

		exec.getter.EXPECT().GetAll(context.Background()).Return(globalURLs, nil)
		stream.EXPECT().Send(resp).Return(errors.New("stream error"))

		err := exec.handler.GetAllURL(req, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream successfully send all response", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetAllURLRequest{}
		stream := mock_grpc.NewMockURLShortenerService_GetAllURLServer(ctrl)

		exec.getter.EXPECT().GetAll(context.Background()).Return(globalURLs, nil)
		stream.EXPECT().Send(globalsResponses[0]).Return(nil)
		stream.EXPECT().Send(globalsResponses[1]).Return(nil)
		stream.EXPECT().Send(globalsResponses[2]).Return(nil)

		err := exec.handler.GetAllURL(req, stream)

		assert.Nil(t, err)
	})
}

func TestURLShortener_GetURLDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)

		resp, err := exec.handler.GetURLDetail(context.Background(), nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL, err)
		assert.Nil(t, resp)
	})

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetURLDetailRequest{ShortUrl: "http://short-1.url"}
		exec.getter.EXPECT().GetByShortURL(context.Background(), req.GetShortUrl()).Return(nil, entity.ErrInternalServer)

		resp, err := exec.handler.GetURLDetail(context.Background(), req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Nil(t, resp)
	})

	t.Run("url can't be found", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetURLDetailRequest{ShortUrl: "http://short-1.url"}
		exec.getter.EXPECT().GetByShortURL(context.Background(), req.GetShortUrl()).Return(nil, entity.ErrURLNotFound)

		resp, err := exec.handler.GetURLDetail(context.Background(), req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrURLNotFound, err)
		assert.Nil(t, resp)
	})

	t.Run("successfully get a single url", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		req := &shortenerv1.GetURLDetailRequest{ShortUrl: "http://short-1.url"}
		exec.getter.EXPECT().GetByShortURL(context.Background(), req.GetShortUrl()).Return(globalURLs[0], nil)

		resp, err := exec.handler.GetURLDetail(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "http://short-1.url", resp.GetShortUrl())
		assert.Equal(t, "http://original-1.url", resp.GetOriginalUrl())
		assert.Equal(t, timestamppb.New(globalURLs[0].ExpiredAt), resp.GetExpiredAt())
	})
}

func createGetAllURLReponse(urls []*entity.URL) []*shortenerv1.GetAllURLResponse {
	result := []*shortenerv1.GetAllURLResponse{}
	for _, url := range urls {
		tmp := &shortenerv1.GetAllURLResponse{
			ShortUrl:    url.ShortURL,
			OriginalUrl: url.OriginalURL,
			ExpiredAt:   timestamppb.New(url.ExpiredAt),
		}
		result = append(result, tmp)
	}
	return result
}

func createURLShortenerExecutor(ctrl *gomock.Controller) *URLShortenerExecutor {
	c := mock_usecase.NewMockCreateShortURL(ctrl)
	g := mock_usecase.NewMockGetURL(ctrl)
	h := handler.NewURLShortener(c, g)
	return &URLShortenerExecutor{
		handler: h,
		creator: c,
		getter:  g,
	}
}
