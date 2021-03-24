package builder

import (
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	"github.com/indrasaputra/url-shortener/internal/repository"
	"github.com/indrasaputra/url-shortener/internal/tool"
	"github.com/indrasaputra/url-shortener/usecase"
)

// BuildGRPCShortURLCreator build ShortURLCreator handler together with all of its dependencies.
func BuildGRPCShortURLCreator(domain string) *handler.ShortURLCreator {
	gen := tool.NewShortURLGenerator(domain)
	repo := repository.NewInMemoryURLRepository()
	creator := usecase.NewShortURLCreator(gen, repo)
	return handler.NewShortURLCreator(creator)
}
