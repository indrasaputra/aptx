package repository

import (
	"context"

	"github.com/indrasaputra/aptx/entity"
)

// GetURLDatabase defines the interface to get URL from the database.
type GetURLDatabase interface {
	// GetAll gets all URLs from the database.
	GetAll(ctx context.Context) ([]*entity.URL, error)
	// GetByCode gets a single URL from the database.
	// If the URL can't be found, it returns entity.ErrNotFound().
	GetByCode(ctx context.Context, code string) (*entity.URL, error)
}

// GetURLCache defines the interface to get URL from the cache.
type GetURLCache interface {
	// Get gets a URL from the cache.
	Get(ctx context.Context, code string) (*entity.URL, error)
	InsertURLCache
}

// URLGetter is responsible to get URL from storage.
// It uses database and cache.
type URLGetter struct {
	database GetURLDatabase
	cache    GetURLCache
}

// NewURLGetter creates an instance of URLGetter.
func NewURLGetter(database GetURLDatabase, cache GetURLCache) *URLGetter {
	return &URLGetter{
		database: database,
		cache:    cache,
	}
}

// GetAll gets all URLs from storage.
func (ug *URLGetter) GetAll(ctx context.Context) ([]*entity.URL, error) {
	return ug.database.GetAll(ctx)
}

// GetByCode gets a single URL from storage.
// If the URL can't be found, it returns entity.ErrNotFound().
func (ug *URLGetter) GetByCode(ctx context.Context, code string) (*entity.URL, error) {
	url, cerr := ug.cache.Get(ctx, code)
	if cerr != nil {
		return nil, cerr
	}
	if url != nil {
		return url, nil
	}

	url, derr := ug.database.GetByCode(ctx, code)
	if derr != nil {
		return nil, derr
	}
	_ = ug.cache.Save(ctx, url)
	return url, nil
}
