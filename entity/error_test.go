package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/entity"
)

func TestErrInternal(t *testing.T) {
	t.Run("success get codes.Internal error", func(t *testing.T) {
		err := entity.ErrInternal("")

		assert.Contains(t, err.Error(), "rpc error: code = Internal")
	})
}

func TestErrEmptyURL(t *testing.T) {
	t.Run("success get codes.InvalidArgument error", func(t *testing.T) {
		err := entity.ErrEmptyURL()

		assert.Contains(t, err.Error(), "rpc error: code = InvalidArgument")
	})
}

func TestErrAlreadyExists(t *testing.T) {
	t.Run("success get codes.AlreadyExists error", func(t *testing.T) {
		err := entity.ErrAlreadyExists()

		assert.Contains(t, err.Error(), "rpc error: code = AlreadyExists")
	})
}

func TestErrNotFound(t *testing.T) {
	t.Run("success get codes.NotFound error", func(t *testing.T) {
		err := entity.ErrNotFound()

		assert.Contains(t, err.Error(), "rpc error: code = NotFound")
	})
}
