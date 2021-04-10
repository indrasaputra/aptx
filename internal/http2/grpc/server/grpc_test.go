package server_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/indrasaputra/url-shortener/internal/http2/grpc/server"
)

var (
	testGRPCPort = "8080"
)

func TestNewGRPC(t *testing.T) {
	t.Run("successfully create a gRPC server", func(t *testing.T) {
		srv := server.NewGRPC(testGRPCPort)
		assert.NotNil(t, srv)
	})
}

func TestGRPC_Run(t *testing.T) {
	t.Run("listener fails", func(t *testing.T) {
		srv := server.NewGRPC("abc")

		err := srv.Run()
		defer srv.Stop()

		assert.NotNil(t, err)
	})

	t.Run("success run", func(t *testing.T) {
		srv := server.NewGRPC("8018")

		err := srv.Run()
		defer srv.Stop()
		time.Sleep(1 * time.Second)

		assert.Nil(t, err)
	})
}

func TestGRPC_RegisterServices(t *testing.T) {
	t.Run("success register service", func(t *testing.T) {
		srv := server.NewGRPC(testGRPCPort)
		fn := func(server *grpc.Server) {}

		assert.NotPanics(t, func() { srv.RegisterServices(fn) })
	})
}
