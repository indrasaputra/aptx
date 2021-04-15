package repository_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/aptx/entity"
	"github.com/indrasaputra/aptx/internal/repository"
	mock_repository "github.com/indrasaputra/aptx/test/mock/repository"
)

type URLGetterExecutor struct {
	getter   *repository.URLGetter
	database *mock_repository.MockGetURLDatabase
	cache    *mock_repository.MockGetURLCache
}

func TestNewURLGetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of URLGetter", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		assert.NotNil(t, exec.getter)
	})
}

func TestURLGetter_GetByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("cache returns error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.cache.EXPECT().Get(context.Background(), testURL.Code).Return(nil, entity.ErrInternal(""))

		res, err := exec.getter.GetByCode(context.Background(), testURL.Code)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("url found in cache", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.cache.EXPECT().Get(context.Background(), testURL.Code).Return(testURL, nil)

		res, err := exec.getter.GetByCode(context.Background(), testURL.Code)

		assert.Nil(t, err)
		assert.Equal(t, testURL, res)
	})

	t.Run("database returns error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.cache.EXPECT().Get(context.Background(), testURL.Code).Return(nil, nil)
		exec.database.EXPECT().GetByCode(context.Background(), testURL.Code).Return(nil, entity.ErrInternal(""))

		res, err := exec.getter.GetByCode(context.Background(), testURL.Code)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("success get url from db and save to cache", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.cache.EXPECT().Get(context.Background(), testURL.Code).Return(nil, nil)
		exec.database.EXPECT().GetByCode(context.Background(), testURL.Code).Return(testURL, nil)
		exec.cache.EXPECT().Save(context.Background(), testURL).Return(nil)

		res, err := exec.getter.GetByCode(context.Background(), testURL.Code)

		assert.Nil(t, err)
		assert.Equal(t, testURL, res)
	})
}

func TestURLGetter_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("database returns error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.database.EXPECT().GetAll(context.Background()).Return([]*entity.URL{}, entity.ErrInternal(""))

		res, err := exec.getter.GetAll(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Empty(t, res)
	})

	t.Run("database returns empty list and nil error", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.database.EXPECT().GetAll(context.Background()).Return([]*entity.URL{}, nil)

		res, err := exec.getter.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Empty(t, res)
	})

	t.Run("success get url from db", func(t *testing.T) {
		exec := createURLGetterExecutor(ctrl)
		exec.database.EXPECT().GetAll(context.Background()).Return([]*entity.URL{testURL}, nil)

		res, err := exec.getter.GetAll(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, testURL, res)
	})
}

func createURLGetterExecutor(ctrl *gomock.Controller) *URLGetterExecutor {
	d := mock_repository.NewMockGetURLDatabase(ctrl)
	c := mock_repository.NewMockGetURLCache(ctrl)
	i := repository.NewURLGetter(d, c)
	return &URLGetterExecutor{
		getter:   i,
		database: d,
		cache:    c,
	}
}
