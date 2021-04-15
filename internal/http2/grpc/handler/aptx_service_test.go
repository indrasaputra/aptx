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
	testShortenURLRequest    = &aptxv1.ShortenURLRequest{OriginalUrl: testOriginalURL}
	testShortenURLResponse   = &aptxv1.ShortenURLResponse{Url: testShortenerV1URL}
	testGetAllURLRequest     = &aptxv1.GetAllURLRequest{}
	testGetAllURLResponse    = &aptxv1.GetAllURLResponse{Urls: []*aptxv1.URL{testShortenerV1URL}}
	testStreamAllURLRequest  = &aptxv1.StreamAllURLRequest{}
	testStreamAllURLResponse = &aptxv1.StreamAllURLResponse{Url: testShortenerV1URL}
	testGetURLDetailRequest  = &aptxv1.GetURLDetailRequest{Code: testCode}
	testGetURLDetailResponse = &aptxv1.GetURLDetailResponse{Url: testShortenerV1URL}
	testURLs                 = []*entity.URL{
		{
			Code:        testCode,
			ShortURL:    testShortURL,
			OriginalURL: testOriginalURL,
			ExpiredAt:   testExpiredAt,
			CreatedAt:   testCreatedAt,
		},
	}
)

type AptxServiceExecutor struct {
	handler *handler.AptxService
	creator *mock_usecase.MockCreateShortURL
	getter  *mock_usecase.MockGetURL
}

func TestNewAptxService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of AptxService", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestAptxService_ShortenURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)

		resp, err := exec.handler.ShortenURL(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("creator usecase returns error", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.creator.EXPECT().Create(testContext, testShortenURLRequest.GetOriginalUrl()).Return(nil, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.ShortenURL(testContext, testShortenURLRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Nil(t, resp)
	})

	t.Run("successfully create a short URL", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.creator.EXPECT().Create(testContext, testShortenURLRequest.GetOriginalUrl()).Return(testURLs[0], nil)

		resp, err := exec.handler.ShortenURL(testContext, testShortenURLRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testShortenURLResponse, resp)
	})
}

func TestAptxService_GetAllURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty request is prohibited", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)

		resp, err := exec.handler.GetAllURL(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.getter.EXPECT().GetAll(testContext).Return([]*entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.GetAllURL(testContext, testGetAllURLRequest)

		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

	t.Run("success get all urls", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)

		resp, err := exec.handler.GetAllURL(testContext, testGetAllURLRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testGetAllURLResponse, resp)
		assert.Equal(t, 1, len(resp.GetUrls()))
	})
}

func TestAptxService_StreamAllURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		stream := mock_grpc.NewMockAptxService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return([]*entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream can't send response", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		stream := mock_grpc.NewMockAptxService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)
		stream.EXPECT().Send(testStreamAllURLResponse).Return(errors.New("stream error"))

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.NotNil(t, err)
	})

	t.Run("stream successfully send all response", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		stream := mock_grpc.NewMockAptxService_StreamAllURLServer(ctrl)
		stream.EXPECT().Context().Return(testContext)
		exec.getter.EXPECT().GetAll(testContext).Return(testURLs, nil)
		stream.EXPECT().Send(testStreamAllURLResponse).Return(nil)

		err := exec.handler.StreamAllURL(testStreamAllURLRequest, stream)

		assert.Nil(t, err)
	})
}

func TestAptxService_GetURLDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty url is prohibited", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)

		resp, err := exec.handler.GetURLDetail(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
		assert.Nil(t, resp)
	})

	t.Run("getter usecase returns error", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(nil, entity.ErrInternal(testErrInternalMessage))

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Nil(t, resp)
	})

	t.Run("url can't be found", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(nil, entity.ErrNotFound())

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, resp)
	})

	t.Run("successfully get a single url", func(t *testing.T) {
		exec := createAptxServiceExecutor(ctrl)
		exec.getter.EXPECT().GetByCode(testContext, testGetURLDetailRequest.GetCode()).Return(testURLs[0], nil)

		resp, err := exec.handler.GetURLDetail(testContext, testGetURLDetailRequest)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testGetURLDetailResponse, resp)
	})
}

func createAptxServiceExecutor(ctrl *gomock.Controller) *AptxServiceExecutor {
	c := mock_usecase.NewMockCreateShortURL(ctrl)
	g := mock_usecase.NewMockGetURL(ctrl)
	h := handler.NewAptxService(c, g)
	return &AptxServiceExecutor{
		handler: h,
		creator: c,
		getter:  g,
	}
}
