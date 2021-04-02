package usecase

import (
	"context"

	"github.com/indrasaputra/url-shortener/entity"
)

// GetURL is the interface that defines the URL retrieve process.
type GetURL interface {
	// GetAll gets all URL available in system.
	GetAll(ctx context.Context) ([]*entity.URL, error)
	// GetByCode gets a single data available in system based on code.
	GetByCode(ctx context.Context, code string) (*entity.URL, error)
}

// GetURLRepository defines the contract to get URL.
type GetURLRepository interface {
	// GetAll gets all URL available in repository.
	GetAll(ctx context.Context) ([]*entity.URL, error)
	// GetByCode gets a single data available in repository based on code.
	// If the data is not found, it returns ErrURLNotFound.
	GetByCode(ctx context.Context, code string) (*entity.URL, error)
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
func (ug *URLGetter) GetAll(ctx context.Context) ([]*entity.URL, error) {
	return ug.repo.GetAll(ctx)
}

// GetByCode gets a single data available in system based on code.
func (ug *URLGetter) GetByCode(ctx context.Context, code string) (*entity.URL, error) {
	return ug.repo.GetByCode(ctx, code)
}
