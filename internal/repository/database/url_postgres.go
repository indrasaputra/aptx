package database

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/indrasaputra/url-shortener/entity"
)

const (
	errDuplicateMessage = "pq: duplicate key value violates unique constraint"
)

// URLPostgres is responsible to connect URL with PostgreSQL.
type URLPostgres struct {
	db *sql.DB
}

// NewURLPostgres creates an instance of URLPostgres.
func NewURLPostgres(db *sql.DB) *URLPostgres {
	return &URLPostgres{db: db}
}

// Insert inserts a URL into PostgreSQL by running SQL INSERT query.
// It validates if the URL already exists.
func (ur *URLPostgres) Insert(ctx context.Context, url *entity.URL) error {
	if url == nil {
		return entity.ErrEmptyURL()
	}

	query := "INSERT INTO " +
		"urls (code, short_url, original_url, expired_at, created_at) " +
		"VALUES ($1, $2, $3, $4, $5)"

	_, err := ur.db.ExecContext(ctx, query,
		url.Code,
		url.ShortURL,
		url.OriginalURL,
		url.ExpiredAt,
		url.CreatedAt,
	)
	if err != nil && strings.Contains(err.Error(), errDuplicateMessage) {
		return entity.ErrAlreadyExists()
	}
	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	return nil
}

// GetAll gets all URLs from PostgreSQL.
// If there isn't any data, it returns empty list and nil error.
func (ur *URLPostgres) GetAll(ctx context.Context) ([]*entity.URL, error) {
	query := "SELECT code, short_url, original_url, expired_at, created_at FROM urls"
	rows, qerr := ur.db.QueryContext(ctx, query)
	if qerr != nil {
		return []*entity.URL{}, entity.ErrInternal(qerr.Error())
	}
	defer rows.Close()

	res := []*entity.URL{}
	for rows.Next() {
		var tmp entity.URL
		if serr := rows.Scan(&tmp.Code, &tmp.ShortURL, &tmp.OriginalURL, &tmp.ExpiredAt, &tmp.CreatedAt); serr != nil {
			log.Printf("[URLPostgres-GetAll]scan rows error: %s", serr.Error())
			continue
		}
		res = append(res, &tmp)
	}
	if rows.Err() != nil {
		return []*entity.URL{}, entity.ErrInternal(rows.Err().Error())
	}
	return res, nil
}

// GetByCode gets a single URL by its code.
// It returns entity.ErrNotFound() if the URL can't be found.
func (ur *URLPostgres) GetByCode(ctx context.Context, code string) (*entity.URL, error) {
	query := "SELECT code, short_url, original_url, expired_at, created_at FROM urls WHERE code = $1 LIMIT 1"
	row := ur.db.QueryRowContext(ctx, query, code)

	res := entity.URL{}
	err := row.Scan(&res.Code, &res.ShortURL, &res.OriginalURL, &res.ExpiredAt, &res.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, entity.ErrNotFound()
	}
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	return &res, nil
}
