package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

const (
	connProtocol = "tcp"
)

// RegisterServiceFunc defines function contract to register service.
type RegisterServiceFunc func(server *grpc.Server)

// GRPC is responsible to act as gRPC server.
// It composes grpc.Server.
type GRPC struct {
	*grpc.Server
	listener net.Listener
	port     string
}

// NewGRPC creates an instance of GRPC.
func NewGRPC(port string, options ...grpc.ServerOption) *GRPC {
	srv := grpc.NewServer(options...)
	return &GRPC{
		Server: srv,
		port:   port,
	}
}

// Run runs the server.
// It basically runs grpc.Server.Serve in a goroutine.
// So, it is not blocking.
func (g *GRPC) Run() error {
	var err error
	g.listener, err = net.Listen(connProtocol, fmt.Sprintf(":%s", g.port))
	if err != nil {
		return err
	}

	go g.serve()
	return nil
}

// AwaitTerminations blocks the server and wait for termination signal.
// The termination signal must be one of SIGINT or SIGTERM.
// Once it receives one of those signals, the gRPC server will perform graceful stop and close the listener.
func (g *GRPC) AwaitTermination() error {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	g.GracefulStop()
	return g.listener.Close()
}

// RegisterServices registers gRPC service to gRPC server.
func (g *GRPC) RegisterServices(fns ...RegisterServiceFunc) {
	for _, fn := range fns {
		fn(g.Server)
	}
}

func (g *GRPC) serve() {
	if err := g.Serve(g.listener); err != nil {
		panic(err)
	}
}
