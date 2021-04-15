package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/aptx/entity"
	aptxv1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
	"github.com/indrasaputra/aptx/usecase"
)

// URLShortener handles HTTP/2 gRPC request for URL aptx.
type URLShortener struct {
	aptxv1.UnimplementedURLShortenerServiceServer
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
func (us *URLShortener) CreateShortURL(ctx context.Context, request *aptxv1.CreateShortURLRequest) (*aptxv1.CreateShortURLResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyURL()
	}

	url, cerr := us.creator.Create(ctx, request.GetOriginalUrl())
	if cerr != nil {
		return nil, cerr
	}

	return createCreateShortURLResponseFromEntity(url), nil
}

// GetAllURL handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// Its specific job is to get all available URLs in the system.
func (us *URLShortener) GetAllURL(ctx context.Context, request *aptxv1.GetAllURLRequest) (*aptxv1.GetAllURLResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyURL()
	}

	urls, err := us.getter.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	return createGetAllURLResponseFromEntity(urls), nil
}

// StreamAllURL handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// Its specific job is to get all available URLs in the system using stream.
func (us *URLShortener) StreamAllURL(request *aptxv1.StreamAllURLRequest, stream aptxv1.URLShortenerService_StreamAllURLServer) error {
	urls, err := us.getter.GetAll(stream.Context())
	if err != nil {
		return err
	}

	for _, url := range urls {
		resp := createStreamAllURLResponseFromEntity(url)
		if serr := stream.Send(resp); serr != nil {
			return entity.ErrInternal(serr.Error())
		}
	}
	return nil
}

// GetURLDetail handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// Its specific job is to get a detail of a single short URL.
func (us *URLShortener) GetURLDetail(ctx context.Context, request *aptxv1.GetURLDetailRequest) (*aptxv1.GetURLDetailResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyURL()
	}

	url, err := us.getter.GetByCode(ctx, request.GetCode())
	if err != nil {
		return nil, err
	}
	return createGetURLDetailResponseFromEntity(url), nil
}

func createCreateShortURLResponseFromEntity(url *entity.URL) *aptxv1.CreateShortURLResponse {
	return &aptxv1.CreateShortURLResponse{
		Url: createShortenerV1URL(url),
	}
}

func createGetAllURLResponseFromEntity(urls []*entity.URL) *aptxv1.GetAllURLResponse {
	res := &aptxv1.GetAllURLResponse{}
	for _, url := range urls {
		res.Urls = append(res.Urls, createShortenerV1URL(url))
	}
	return res
}

func createStreamAllURLResponseFromEntity(url *entity.URL) *aptxv1.StreamAllURLResponse {
	return &aptxv1.StreamAllURLResponse{
		Url: createShortenerV1URL(url),
	}
}

func createGetURLDetailResponseFromEntity(url *entity.URL) *aptxv1.GetURLDetailResponse {
	return &aptxv1.GetURLDetailResponse{
		Url: createShortenerV1URL(url),
	}
}

func createShortenerV1URL(url *entity.URL) *aptxv1.URL {
	return &aptxv1.URL{
		Code:        url.Code,
		ShortUrl:    url.ShortURL,
		OriginalUrl: url.OriginalURL,
		ExpiredAt:   timestamppb.New(url.ExpiredAt),
		CreatedAt:   timestamppb.New(url.CreatedAt),
	}
}
