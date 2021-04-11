package database_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/indrasaputra/url-shortener/entity"
	"github.com/indrasaputra/url-shortener/internal/repository/database"
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
	testExpiredAtString = "time.Now().Add(1 * time.Minute)"
	testCreatedAtString = "time.Now()"
	errDuplicate        = errors.New("pq: duplicate key value violates unique constraint")
	errPostgresInternal = errors.New("database down")
)

type URLPostgresExecutor struct {
	database *database.URLPostgres
	pgx      pgxmock.PgxPoolIface
}

func TestNewURLPostgres(t *testing.T) {
	t.Run("successfully create an instance of URLPostgres", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		assert.NotNil(t, exec.database)
	})
}

func TestURLPostgres_Insert(t *testing.T) {
	t.Run("nil URL is prohibited", func(t *testing.T) {
		exec := createURLPostgresExecutor()

		err := exec.database.Insert(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyURL(), err)
	})

	t.Run("can't insert duplicated data", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.ExpectExec(`INSERT INTO urls \(code, short_url, original_url, expired_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WillReturnError(errDuplicate)

		err := exec.database.Insert(testContext, testURL)

		assert.NotNil(t, err)
		assert.Equal(t, codes.AlreadyExists, status.Code(err))
	})

	t.Run("postgres returns internal error", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectExec(`INSERT INTO urls \(code, short_url, original_url, expired_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WillReturnError(errPostgresInternal)

		err := exec.database.Insert(testContext, testURL)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success insert a new URL", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectExec(`INSERT INTO urls \(code, short_url, original_url, expired_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err := exec.database.Insert(testContext, testURL)

		assert.Nil(t, err)
	})
}

func TestURLPostgres_GetAll(t *testing.T) {
	t.Run("select all query returns error", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls`).
			WillReturnError(errPostgresInternal)

		res, err := exec.database.GetAll(testContext)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Empty(t, res)
	})

	t.Run("select all rows scan returns error", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls`).
			WillReturnRows(pgxmock.
				NewRows([]string{"code", "short_url", "original_url", "expired_at", "created_at"}).
				AddRow(testURLCode, testURLShort, testURLOriginal, testURLExpiredAt, testURLCreatedAt).
				AddRow(testURLCode, testURLShort, testURLOriginal, testExpiredAtString, testCreatedAtString),
			)

		res, err := exec.database.GetAll(testContext)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
	})

	t.Run("select all rows error occurs after scanning", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls`).
			WillReturnRows(pgxmock.
				NewRows([]string{"code", "short_url", "original_url", "expired_at", "created_at"}).
				AddRow(testURLCode, testURLShort, testURLOriginal, testURLExpiredAt, testURLCreatedAt).
				AddRow(testURLCode, testURLShort, testURLOriginal, testExpiredAtString, testCreatedAtString).
				RowError(2, errPostgresInternal),
			)

		res, err := exec.database.GetAll(testContext)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Empty(t, res)
	})

	t.Run("successfully retrieve all rows", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls`).
			WillReturnRows(pgxmock.
				NewRows([]string{"code", "short_url", "original_url", "expired_at", "created_at"}).
				AddRow(testURLCode, testURLShort, testURLOriginal, testURLExpiredAt, testURLCreatedAt).
				AddRow(testURLCode, testURLShort, testURLOriginal, testURLExpiredAt, testURLCreatedAt),
			)

		res, err := exec.database.GetAll(testContext)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(res))
	})
}

func TestURLPostgres_GetByCode(t *testing.T) {
	t.Run("select by code query returns empty row", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls WHERE code = \$1 LIMIT 1`).
			WillReturnError(pgx.ErrNoRows)

		res, err := exec.database.GetByCode(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, res)
	})

	t.Run("select by code query returns error", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls WHERE code = \$1 LIMIT 1`).
			WillReturnError(errPostgresInternal)

		res, err := exec.database.GetByCode(testContext, testURLCode)

		assert.NotNil(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Nil(t, res)
	})

	t.Run("successfully retrieve row", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.
			ExpectQuery(`SELECT code, short_url, original_url, expired_at, created_at FROM urls WHERE code = \$1 LIMIT 1`).
			WillReturnRows(pgxmock.
				NewRows([]string{"code", "short_url", "original_url", "expired_at", "created_at"}).
				AddRow(testURLCode, testURLShort, testURLOriginal, testURLExpiredAt, testURLCreatedAt),
			)

		res, err := exec.database.GetByCode(testContext, testURLCode)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestURLPostgres_IsAlive(t *testing.T) {
	t.Run("postgres is not alive", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.ExpectPing().WillReturnError(errPostgresInternal)

		alive := exec.database.IsAlive(testContext)

		assert.False(t, alive)
	})

	t.Run("postgres is not alive", func(t *testing.T) {
		exec := createURLPostgresExecutor()
		exec.pgx.ExpectPing().WillReturnError(nil)

		alive := exec.database.IsAlive(testContext)

		assert.True(t, alive)
	})
}

func createURLPostgresExecutor() *URLPostgresExecutor {
	mock, err := pgxmock.NewPool(pgxmock.MonitorPingsOption(true))
	if err != nil {
		log.Panicf("error opening a stub database connection: %v\n", err)
	}

	database := database.NewURLPostgres(mock)
	return &URLPostgresExecutor{
		database: database,
		pgx:      mock,
	}
}
