package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/indrasaputra/shortener/entity"
)

const (
	maxRetry          = 3
	defaultExpiryTime = 7 * 24 * time.Hour
	shortURLLength    = 7
)

// CreateShortURL is the interface that defines the short url creation.
type CreateShortURL interface {
	// Create creates a short version of the given URL.
	Create(ctx context.Context, url string) (*entity.URL, *entity.Error)
}

// URLGenerator defines the short url generator.
type URLGenerator interface {
	// Generate generates a short URL with defined length.
	Generate(length uint) (string, *entity.Error)
}

// URLRepository defines the repository for URL.
type URLRepository interface {
	// Save saves the URL in the repository.
	Save(ctx context.Context, url *entity.URL) *entity.Error
}

// ShortURLCreator is responsible for creating a unique short URL.
type ShortURLCreator struct {
	generator URLGenerator
	repo      URLRepository
}

// NewShortURLCreator creates an instance of ShortURLCreator.
func NewShortURLCreator(generator URLGenerator, repo URLRepository) *ShortURLCreator {
	return &ShortURLCreator{
		generator: generator,
		repo:      repo,
	}
}

// Create creates a short URL for the given URL.
// It tries to ensure that the short URL is unique from the rest.
// If it is unsuccessful in creating a unique short URL, it will return error.
//
// Currently, it does not check if the URL is valid. It only checks whether the URL is empty.
func (sc *ShortURLCreator) Create(ctx context.Context, url string) (*entity.URL, *entity.Error) {
	if strings.TrimSpace(url) == "" {
		return nil, entity.ErrEmptyURL
	}

	var data *entity.URL
	var err *entity.Error
	for i := 0; i < maxRetry; i++ {
		data, err = sc.generateURL(url)
		if err != nil {
			continue
		}
		if err = sc.repo.Save(ctx, data); err == nil {
			return data, nil
		}
	}
	return nil, err
}

func (sc *ShortURLCreator) generateURL(url string) (*entity.URL, *entity.Error) {
	shortURL, err := sc.generator.Generate(shortURLLength)
	if err != nil {
		return nil, err
	}
	return &entity.URL{
		ShortURL:    shortURL,
		OriginalURL: url,
		ExpiredAt:   time.Now().Add(defaultExpiryTime),
	}, nil
}
