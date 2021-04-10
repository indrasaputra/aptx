package server_test

import (
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/internal/http2/grpc/server"
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
		fn := func(server *runtime.ServeMux) {}
		assert.NotPanics(t, func() { srv.RegisterEndpoints(fn) })
	})
}
