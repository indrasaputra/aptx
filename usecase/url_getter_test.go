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
	testURLs = []*entity.URL{
		{
			Code:        testURLCode,
			ShortURL:    testURLShort,
			OriginalURL: testURLOriginal,
			ExpiredAt:   time.Now().Add(1 * time.Minute),
			CreatedAt:   time.Now(),
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
		exec.repo.EXPECT().GetAll(context.Background()).Return([]*entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		urls, err := exec.usecase.GetAll(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Empty(t, urls)
	})

	t.Run("success get all urls", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(context.Background()).Return(testURLs, nil)

		urls, err := exec.usecase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, testURLs, urls)
	})
}

func TestURLGetter_GetByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("repository returns error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetByCode(context.Background(), testURLCode).Return(&entity.URL{}, entity.ErrInternal(testErrInternalMessage))

		urls, err := exec.usecase.GetByCode(context.Background(), testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(testErrInternalMessage), err)
		assert.Empty(t, urls)
	})

	t.Run("url can't be found", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetByCode(context.Background(), testURLCode).Return(&entity.URL{}, entity.ErrNotFound())

		urls, err := exec.usecase.GetByCode(context.Background(), testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Empty(t, urls)
	})

	t.Run("success get single url", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.repo.EXPECT().GetByCode(context.Background(), testURLCode).Return(testURLs[0], nil)

		urls, err := exec.usecase.GetByCode(context.Background(), testURLCode)

		assert.Nil(t, err)
		assert.Equal(t, testURLs[0], urls)
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
