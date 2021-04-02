package builder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/internal/builder"
)

func TestBuildGRPCURLShortener(t *testing.T) {
	t.Run("successfully create URLShortener handler", func(t *testing.T) {
		hdr := builder.BuildGRPCURLShortener("http://short-url.com")
		assert.NotNil(t, hdr)
	})
}

func TestBuildGRPCHealthChecker(t *testing.T) {
	t.Run("successfully create HealthChecker handler", func(t *testing.T) {
		hdr := builder.BuildGRPCHealthChecker()
		assert.NotNil(t, hdr)
	})
}
