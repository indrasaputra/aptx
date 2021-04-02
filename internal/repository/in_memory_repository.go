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
// It checks whether the code / short URL already exists in the storage.
// If the data already exists, it will return entity.ErrAlreadyExists().
func (ir *InMemoryURLRepository) Save(ctx context.Context, url *entity.URL) error {
	_, exist := ir.data[url.Code]
	if exist {
		return entity.ErrAlreadyExists()
	}

	ir.data[url.Code] = url
	return nil
}

// GetAll gets all URLs in storage.
// Since the implementation uses HashMap, the data may be unordered.
func (ir *InMemoryURLRepository) GetAll(ctx context.Context) ([]*entity.URL, error) {
	urls := []*entity.URL{}
	for _, url := range ir.data {
		urls = append(urls, url)
	}
	return urls, nil
}

// GetByCode gets a single URLs in storage.
// If the URL can't be found, it returns entity.ErrNotFound().
func (ir *InMemoryURLRepository) GetByCode(ctx context.Context, code string) (*entity.URL, error) {
	url, found := ir.data[code]
	if !found {
		return nil, entity.ErrNotFound()
	}
	return url, nil
}

// IsAlive always returns true since the HashMap is always alive inside the system.
func (ir *InMemoryURLRepository) IsAlive(_ context.Context) bool {
	return true
}
