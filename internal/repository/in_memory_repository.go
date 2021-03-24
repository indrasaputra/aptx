package repository

import (
	"context"

	"github.com/indrasaputra/url-shortener/entity"
)

// InMemoryURLRepository uses memory as storage.
// It uses a HashMap data structure to store the URL.
type InMemoryURLRepository struct {
	data map[string]*entity.URL
}

// NewInMemoryURLRepository creates an instance of InMemoryURLRepository.
func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		data: make(map[string]*entity.URL),
	}
}

// Save saves a new data into storage.
// It checks whether the short URL already exists in the storage.
// If the data already exists, it will return ErrDuplicatedShortURL.
func (ir *InMemoryURLRepository) Save(ctx context.Context, url *entity.URL) *entity.Error {
	_, exist := ir.data[url.ShortURL]
	if exist {
		return entity.ErrDuplicatedShortURL
	}

	ir.data[url.ShortURL] = url
	return nil
}
