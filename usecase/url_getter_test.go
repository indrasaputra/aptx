package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
	mock_usecase "github.com/indrasaputra/url-shortener/test/mock/usecase"
	"github.com/indrasaputra/url-shortener/usecase"
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
)

type URLGetterExecutor struct {
	usecase *usecase.URLGetter
	repo    *mock_usecase.MockGetURLRepository
}

func TestNewURLGetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of URLGetter", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func TestURLGetter_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("repository returns error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(context.Background()).Return([]*entity.URL{}, entity.ErrInternalServer)

		urls, err := exec.usecase.GetAll(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternalServer, err)
		assert.Empty(t, urls)
	})

	t.Run("success get all urls", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(context.Background()).Return(globalURLs, nil)

		urls, err := exec.usecase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, globalURLs, urls)
	})
}

func createURLGetterExecutor(ctrl *gomock.Controller) *URLGetterExecutor {
	r := mock_usecase.NewMockGetURLRepository(ctrl)
	u := usecase.NewURLGetter(r)

	return &URLGetterExecutor{
		usecase: u,
		repo:    r,
	}
}
