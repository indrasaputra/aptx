package server_test

import (
	"errors"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/aptx/internal/http2/grpc/server"
)

var (
	testRestPort = "8081"
)

func TestNewRest(t *testing.T) {
	t.Run("success create rest server", func(t *testing.T) {
		srv := server.NewRest(testRestPort)
		assert.NotNil(t, srv)
	})
}

func TestRest_EnablePrometheus(t *testing.T) {
	t.Run("success enable prometheus", func(t *testing.T) {
		srv := server.NewRest(testRestPort)
		err := srv.EnablePrometheus()
		assert.Nil(t, err)
	})
}

func TestRest_RegisterEndpoints(t *testing.T) {
	t.Run("success register endpoint", func(t *testing.T) {
		srv := server.NewRest(testRestPort)
		fn := func(server *runtime.ServeMux) error { return errors.New("endpoint doesn't exist") }

		err := srv.RegisterEndpoints(fn)

		assert.NotNil(t, err)
	})

	t.Run("success register endpoint", func(t *testing.T) {
		srv := server.NewRest(testRestPort)
		fn := func(server *runtime.ServeMux) error { return nil }

		err := srv.RegisterEndpoints(fn)

		assert.Nil(t, err)
	})
}
