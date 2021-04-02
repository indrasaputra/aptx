package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpchealthv1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	shortenerHandler := builder.BuildGRPCURLShortener(cfg.Domain)
	healthCheckerHandler := builder.BuildGRPCHealthChecker()

	grpcServer := createGRPCServer()
	registerGRPCServer(grpcServer, shortenerHandler, healthCheckerHandler)
	runGRPCServer(grpcServer, cfg)

	restServer := createRestServer(cfg)
	registerRestServer(restServer)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", cfg.PortHTTP), restServer)
}

func createGRPCServer() *grpc.Server {
	logger, zerr := zap.NewProduction()
	checkError(zerr)

	grpc_zap.ReplaceGrpcLoggerV2(logger)
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	return server
}

func registerGRPCServer(server *grpc.Server, shortenerHandler *handler.URLShortener, healthCheckerHandler *handler.HealthChecker) {
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server)

	shortenerv1.RegisterURLShortenerServiceServer(server, shortenerHandler)
	grpchealthv1.RegisterHealthServer(server, healthCheckerHandler)
}

func runGRPCServer(server *grpc.Server, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.PortGRPC))
	checkError(err)
	go func() {
		_ = server.Serve(lis)
	}()
}

func createRestServer(cfg *config.Config) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	serr := shortenerv1.RegisterURLShortenerServiceHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf(":%s", cfg.PortGRPC), opts)
	checkError(serr)
	return mux
}

func registerRestServer(server *runtime.ServeMux) {
	merr := server.HandlePath(http.MethodGet, "/metrics", promHandler())
	checkError(merr)
}

func promHandler() runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
