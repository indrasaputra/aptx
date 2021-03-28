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

// GetAll gets all URLs in storage.
// Since the implementation uses HashMap, the data may be unordered.
func (ir *InMemoryURLRepository) GetAll(ctx context.Context) ([]*entity.URL, *entity.Error) {
	urls := []*entity.URL{}
	for _, url := range ir.data {
		urls = append(urls, url)
	}
	return urls, nil
}

// GetByShortURL gets a single URLs in storage.
// If the URL can't be found, it returns ErrURLNotFound.
func (ir *InMemoryURLRepository) GetByShortURL(ctx context.Context, shortURL string) (*entity.URL, *entity.Error) {
	url, found := ir.data[shortURL]
	if !found {
		return nil, entity.ErrURLNotFound
	}
	return url, nil
}
