package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/repository"
	mock_repository "github.com/indrasaputra/url-shortener/test/mock/repository"
	"github.com/stretchr/testify/assert"
)

var (
	testContext      = context.Background()
	testURLCode      = "AbCdE12"
	testURLShort     = "http://localhost/" + testURLCode
	testURLOriginal  = "http://very-long-url.url"
	testURLExpiredAt = time.Now().Add(1 * time.Minute)
	testURLCreatedAt = time.Now()
	testURL          = &entity.URL{
		Code:        testURLCode,
		ShortURL:    testURLShort,
		OriginalURL: testURLOriginal,
		ExpiredAt:   testURLExpiredAt,
		CreatedAt:   testURLCreatedAt,
	}
)

type URLInserterExecutor struct {
	creator  *repository.URLInserter
	database *mock_repository.MockInsertURLDatabase
	cache    *mock_repository.MockInsertURLCache
}

func TestNewURLInserter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of URLInserter", func(t *testing.T) {
		exec := createURLInserterExecutor(ctrl)
		assert.NotNil(t, exec.creator)
	})
}

func TestURLInserter_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty toggle is prohibited", func(t *testing.T) {
		exec := createURLInserterExecutor(ctrl)

		err := exec.creator.Save(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
	})

	t.Run("database returns error", func(t *testing.T) {
		exec := createURLInserterExecutor(ctrl)
		exec.database.EXPECT().Insert(testContext, testURL).Return(entity.ErrInternal(""))

		err := exec.creator.Save(testContext, testURL)

		assert.NotNil(t, err)
	})

	t.Run("cache error is ignored", func(t *testing.T) {
		exec := createURLInserterExecutor(ctrl)
		exec.database.EXPECT().Insert(testContext, testURL).Return(nil)
		exec.cache.EXPECT().Save(testContext, testURL).Return(entity.ErrInternal(""))

		err := exec.creator.Save(testContext, testURL)

		assert.Nil(t, err)
	})

	t.Run("all steps are successful", func(t *testing.T) {
		exec := createURLInserterExecutor(ctrl)
		exec.database.EXPECT().Insert(testContext, testURL).Return(nil)
		exec.cache.EXPECT().Save(testContext, testURL).Return(nil)

		err := exec.creator.Save(testContext, testURL)

		assert.Nil(t, err)
	})
}

func createURLInserterExecutor(ctrl *gomock.Controller) *URLInserterExecutor {
	d := mock_repository.NewMockInsertURLDatabase(ctrl)
	c := mock_repository.NewMockInsertURLCache(ctrl)
	i := repository.NewURLInserter(d, c)
	return &URLInserterExecutor{
		creator:  i,
		database: d,
		cache:    c,
	}
}
