package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/shortener/entity"
	shortenerv1 "github.com/indrasaputra/shortener/proto/indrasaputra/shortener/v1"
	"github.com/indrasaputra/shortener/usecase"
)

// ShortURLCreator handles HTTP/2 gRPC request for creating a new short URL.
type ShortURLCreator struct {
	shortenerv1.UnimplementedURLShortenerServiceServer
	creator usecase.CreateShortURL
}

// NewShortURLCreator creates an instance of ShortURLCreator.
func NewShortURLCreator(creator usecase.CreateShortURL) *ShortURLCreator {
	return &ShortURLCreator{
		creator: creator,
	}
}

// CreateShortURL handles HTTP/2 gRPC request similar to POST in HTTP/1.1.
func (sc *ShortURLCreator) CreateShortURL(ctx context.Context, request *shortenerv1.CreateShortURLRequest) (*shortenerv1.CreateShortURLResponse, error) {
	url, cerr := sc.creator.Create(ctx, request.GetOriginalUrl())
	if cerr != nil {
		return nil, cerr
	}

	return convertURLToCreateShortURLResponse(url), nil
}

func convertURLToCreateShortURLResponse(url *entity.URL) *shortenerv1.CreateShortURLResponse {
	return &shortenerv1.CreateShortURLResponse{
		ShortUrl:  url.ShortURL,
		ExpiredAt: timestamppb.New(url.ExpiredAt),
	}
}
