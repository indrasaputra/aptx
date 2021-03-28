package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/url-shortener/entity"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
	"github.com/indrasaputra/url-shortener/usecase"
)

// URLShortener handles HTTP/2 gRPC request for URL shortener .
type URLShortener struct {
	shortenerv1.UnimplementedURLShortenerServiceServer
	creator usecase.CreateShortURL
	getter  usecase.GetURL
}

// NewURLShortener creates an instance of URLShortener.
func NewURLShortener(creator usecase.CreateShortURL, getter usecase.GetURL) *URLShortener {
	return &URLShortener{
		creator: creator,
		getter:  getter,
	}
}

// CreateShortURL handles HTTP/2 gRPC request similar to POST in HTTP/1.1.
func (us *URLShortener) CreateShortURL(ctx context.Context, request *shortenerv1.CreateShortURLRequest) (*shortenerv1.CreateShortURLResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyURL
	}

	url, cerr := us.creator.Create(ctx, request.GetOriginalUrl())
	if cerr != nil {
		return nil, cerr
	}

	return convertURLToCreateShortURLResponse(url), nil
}

// GetAllURL handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// Its specific job is to get all available URLs in the system.
func (us *URLShortener) GetAllURL(request *shortenerv1.GetAllURLRequest, stream shortenerv1.URLShortenerService_GetAllURLServer) error {
	urls, err := us.getter.GetAll(context.Background())
	if err != nil {
		return err
	}

	for _, url := range urls {
		resp := convertURLToGetAllURLResponse(url)
		if serr := stream.Send(resp); serr != nil {
			return serr
		}
	}
	return nil
}

// GetURLDetail handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// Its specific job is to get a detail of a single short URL.
func (us *URLShortener) GetURLDetail(ctx context.Context, request *shortenerv1.GetURLDetailRequest) (*shortenerv1.GetURLDetailResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyURL
	}

	url, err := us.getter.GetByShortURL(ctx, request.GetShortUrl())
	if err != nil {
		return nil, err
	}
	return convertURLToGetURLDetailResponse(url), nil
}

func convertURLToGetAllURLResponse(url *entity.URL) *shortenerv1.GetAllURLResponse {
	return &shortenerv1.GetAllURLResponse{
		ShortUrl:    url.ShortURL,
		OriginalUrl: url.OriginalURL,
		ExpiredAt:   timestamppb.New(url.ExpiredAt),
	}
}

func convertURLToGetURLDetailResponse(url *entity.URL) *shortenerv1.GetURLDetailResponse {
	return &shortenerv1.GetURLDetailResponse{
		ShortUrl:    url.ShortURL,
		OriginalUrl: url.OriginalURL,
		ExpiredAt:   timestamppb.New(url.ExpiredAt),
	}
}

func convertURLToCreateShortURLResponse(url *entity.URL) *shortenerv1.CreateShortURLResponse {
	return &shortenerv1.CreateShortURLResponse{
		ShortUrl:  url.ShortURL,
		ExpiredAt: timestamppb.New(url.ExpiredAt),
	}
}
