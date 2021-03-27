package builder

import (
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	"github.com/indrasaputra/url-shortener/internal/repository"
	"github.com/indrasaputra/url-shortener/internal/tool"
	"github.com/indrasaputra/url-shortener/usecase"
)

// BuildGRPCURLShortener build URLShortener handler together with all of its dependencies.
func BuildGRPCURLShortener(domain string) *handler.URLShortener {
	gen := tool.NewShortURLGenerator(domain)
	repo := repository.NewInMemoryURLRepository()
	creator := usecase.NewShortURLCreator(gen, repo)
	getter := usecase.NewURLGetter(repo)
	return handler.NewURLShortener(creator, getter)
}
