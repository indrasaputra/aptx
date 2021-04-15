package handler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/aptx/entity"
	"github.com/indrasaputra/aptx/internal/http2/grpc/handler"
	aptxv1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
	mock_grpc "github.com/indrasaputra/aptx/test/mock/http2/grpc"
	mock_usecase "github.com/indrasaputra/aptx/test/mock/usecase"
)

var (
	testErrInternalMessage = "unexpected command"
	testExpiredAt          = time.Now().Add(10 * time.Minute)
	testCreatedAt          = time.Now()
	testContext            = context.Background()
	testCode               = "ABCdef12"
	testShortURL           = "http://short.url/ABCdef12"
	testOriginalURL        = "http://very-long-original.url"
	testShortenerV1URL     = &aptxv1.URL{
		Code:        testCode,
		ShortUrl:    testShortURL,
		OriginalUrl: testOriginalURL,
		ExpiredAt:   timestamppb.New(testExpiredAt),
		CreatedAt:   timestamppb.New(testCreatedAt),
	}
	testCreateShortURLRequest  = &aptxv1.CreateShortURLRequest{OriginalUrl: testOriginalURL}
	testCreateShortURLResponse = &aptxv1.CreateShortURLResponse{Url: testShortenerV1URL}
	testGetAllURLRequest       = &aptxv1.GetAllURLRequest{}
	testGetAllURLResponse      = &aptxv1.GetAllURLResponse{Urls: []*aptxv1.URL{testShortenerV1URL}}
	testStreamAllURLRequest    = &aptxv1.StreamAllURLRequest{}
	testStreamAllURLResponse   = &aptxv1.StreamAllURLResponse{Url: testShortenerV1URL}
	testGetURLDetailRequest    = &aptxv1.GetURLDetailRequest{Code: testCode}
	testGetURLDetailResponse   = &aptxv1.GetURLDetailResponse{Url: testShortenerV1URL}
	testURLs                   = []*entity.URL{
		{
			Code:        testCode,
			ShortURL:    testShortURL,
			OriginalURL: testOriginalURL,
			ExpiredAt:   testExpiredAt,
			CreatedAt:   testCreatedAt,
		},
	}
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

		resp, err := exec.handler.CreateShortURL(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("creator usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.creator.EXPECT().Create(testContext, testCreateShortURLRequest.GetOriginalUrl()).Return(nil, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.CreateShortURL(testContext, testCreateShortURLRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Nil(t, resp)
	})

	t.Run("successfully create a shorturl", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.creator.EXPECT().Create(testContext, testCreateShortURLRequest.GetOriginalUrl()).Return(testURLs[0], nil)

		resp, err := exec.handler.CreateShortURL(testContext, testCreateShortURLRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testCreateShortURLResponse, resp)
	})
}

func TestURLShortener_GetAllURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty request is prohibited", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)

		resp, err := exec.handler.GetAllURL(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.getter.EXPECT().GetAll(testContext).Return([]*entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.GetAllURL(testContext, testGetAllURLRequest)

		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

	t.Run("success get all urls", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)

		resp, err := exec.handler.GetAllURL(testContext, testGetAllURLRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testGetAllURLResponse, resp)
		assert.Equal(t, 1, len(resp.GetUrls()))
	})
}

func TestURLShortener_StreamAllURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		stream := mock_grpc.NewMockURLShortenerService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return([]*entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream can't send response", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		stream := mock_grpc.NewMockURLShortenerService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)
		stream.EXPECT().Send(testStreamAllURLResponse).Return(errors.New("stream error"))

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream successfully send all response", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		stream := mock_grpc.NewMockURLShortenerService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)
		stream.EXPECT().Send(testStreamAllURLResponse).Return(nil)

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.Nil(t, err)
	})
}

func TestURLShortener_GetURLDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)

		resp, err := exec.handler.GetURLDetail(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(nil, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Nil(t, resp)
	})

	t.Run("url can't be found", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(nil, entity.ErrNotFound())

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, resp)
	})

	t.Run("successfully get a single url", func(t *testing.T) {
		exec := createURLShortenerExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(testURLs[0], nil)

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testGetURLDetailResponse, resp)
	})
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
