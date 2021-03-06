package tool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/aptx/internal/tool"
)

func TestNewShortURLGenerator(t *testing.T) {
	t.Run("successfully create an instance of ShortURLGenerator", func(t *testing.T) {
		gen := tool.NewShortURLGenerator("http://aptx.url")
		assert.NotNil(t, gen)
	})
}

func TestShortURLGenerator_Generate(t *testing.T) {
	t.Run("successfully generate random url", func(t *testing.T) {
		domains := []string{"http://localhost///", "http://localhost//", "http://localhost", "http://localhost////"}
		for _, domain := range domains {
			result := make(map[string]bool)
			gen := tool.NewShortURLGenerator(domain)

			// on 10 iterations, the likeliness to produce non-unique string is quite small.
			// if it happenened, then we just got a badluck :P
			for i := 0; i < 10; i++ {
				code, short, err := gen.Generate(7)
				assert.Nil(t, err)

				_, exist := result[code]
				result[code] = true
				assert.False(t, exist)

				assert.Contains(t, short, "http://localhost/")
				assert.Equal(t, len("http://localhost/1234567"), len(short))
				assert.Equal(t, 7, len(code))
			}
		}
	})
}
