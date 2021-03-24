package builder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/internal/builder"
)

func TestBuildGRPCShortURLCreator(t *testing.T) {
	t.Run("successfully create ShortURLCreator handler", func(t *testing.T) {
		hdr := builder.BuildGRPCShortURLCreator("http://short-url.com")
		assert.NotNil(t, hdr)
	})
}
