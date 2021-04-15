package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/indrasaputra/aptx/entity"
)

func TestErrInternal(t *testing.T) {
	t.Run("success get codes.Internal error", func(t *testing.T) {
		err := entity.ErrInternal("")

		assert.Contains(t, err.Error(), "rpc error: code = Internal")
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestErrEmptyURL(t *testing.T) {
	t.Run("success get codes.InvalidArgument error", func(t *testing.T) {
		err := entity.ErrEmptyURL()

		assert.Contains(t, err.Error(), "rpc error: code = InvalidArgument")
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestErrAlreadyExists(t *testing.T) {
	t.Run("success get codes.AlreadyExists error", func(t *testing.T) {
		err := entity.ErrAlreadyExists()

		assert.Contains(t, err.Error(), "rpc error: code = AlreadyExists")
		assert.Equal(t, codes.AlreadyExists, status.Code(err))
	})
}

func TestErrNotFound(t *testing.T) {
	t.Run("success get codes.NotFound error", func(t *testing.T) {
		err := entity.ErrNotFound()

		assert.Contains(t, err.Error(), "rpc error: code = NotFound")
		assert.Equal(t, codes.NotFound, status.Code(err))
	})
}
