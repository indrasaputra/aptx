package repository

import (
	"context"

	"github.com/indrasaputra/url-shortener/entity"
)

// InsertURLDatabase defines the interface to insert a new URL to the database.
type InsertURLDatabase interface {
	// Insert inserts a new URL into the database.
	// It must handle if the data already exists.
	Insert(ctx context.Context, url *entity.URL) error
}

// InsertURLCache defines the interface to insert a new URL to the cache.
type InsertURLCache interface {
	// Save saves a new URL into the cache.
	Save(ctx context.Context, url *entity.URL) error
}

// URLInserter is responsible to insert a new URL into storage.
// It uses database and cache.
type URLInserter struct {
	database InsertURLDatabase
	cache    InsertURLCache
}

// NewURLInserter creates an instance of URLInserter.
func NewURLInserter(database InsertURLDatabase, cache InsertURLCache) *URLInserter {
	return &URLInserter{
		database: database,
		cache:    cache,
	}
}

// Save saves a new URL into the storage.
// First, it inserts to database. If success, the data will be inserted to cache.
// It ignores the error from cache since it can always be generated when retrieving the data.
// But, it doesn't ignore the error from the database.
func (ui *URLInserter) Save(ctx context.Context, url *entity.URL) error {
	if url == nil {
		return entity.ErrEmptyURL()
	}

	if err := ui.database.Insert(ctx, url); err != nil {
		return err
	}
	_ = ui.cache.Save(ctx, url)
	return nil
}
