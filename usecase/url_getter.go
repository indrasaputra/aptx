package usecase

import (
	"context"

	"github.com/indrasaputra/url-shortener/entity"
)

// GetURL is the interface that defines the URL retrieve process.
type GetURL interface {
	// GetAll gets all URL available in system.
	GetAll(ctx context.Context) ([]*entity.URL, *entity.Error)
	// GetByShortURL gets a single data available in system based on short URL.
	GetByShortURL(ctx context.Context, shortURL string) (*entity.URL, *entity.Error)
}

// GetURLRepository defines the contract to get URL.
type GetURLRepository interface {
	// GetAll gets all URL available in repository.
	GetAll(ctx context.Context) ([]*entity.URL, *entity.Error)
	// GetByShortURL gets a single data available in repository based on short URL.
	// If the data is not found, it returns ErrURLNotFound.
	GetByShortURL(ctx context.Context, shortURL string) (*entity.URL, *entity.Error)
}

// URLGetter is responsible to get URL.
type URLGetter struct {
	repo GetURLRepository
}

// NewURLGetter creates an instance of URLGetter.
func NewURLGetter(repo GetURLRepository) *URLGetter {
	return &URLGetter{repo: repo}
}

// GetAll gets all URLs in the system.
func (ug *URLGetter) GetAll(ctx context.Context) ([]*entity.URL, *entity.Error) {
	return ug.repo.GetAll(ctx)
}

// GetByShortURL gets a single data available in system based on short URL.
func (ug *URLGetter) GetByShortURL(ctx context.Context, shortURL string) (*entity.URL, *entity.Error) {
	return ug.repo.GetByShortURL(ctx, shortURL)
}
