package main

import (
	"context"
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/handler"
	"github.com/indrasaputra/url-shortener/internal/http2/grpc/server"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	postgres, perr := builder.BuildPostgresClient(cfg.Postgres)
	checkError(perr)
	redis, rerr := builder.BuildRedisClient(cfg.Redis)
	checkError(rerr)

	shortener := builder.BuildGRPCURLShortener(postgres, redis, cfg.Domain)
	health := builder.BuildGRPCHealthChecker(postgres, redis)

	logger, zerr := zap.NewProduction()
	checkError(zerr)
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	grpcServer := server.NewGRPC(cfg.PortGRPC,
		grpc_middleware.WithUnaryServerChain(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	grpcServer.RegisterServices(
		registerGRPCPrometheus(),
		registerGRPCURLShortenerService(shortener),
		registerGRPCHealthService(health),
	)
	_ = grpcServer.Run()

	restServer := server.NewRest(cfg.PortHTTP)
	promerr := restServer.EnablePrometheus()
	checkError(promerr)

	restServer.RegisterEndpoints(
		registerRestURLShortenerEndpoint(cfg.PortGRPC, grpc.WithInsecure()),
	)
	_ = restServer.Run()

	_ = grpcServer.AwaitTermination()
}

func registerGRPCPrometheus() server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		grpc_prometheus.EnableHandlingTimeHistogram()
		grpc_prometheus.Register(server)
	}
}

func registerGRPCURLShortenerService(shortener *handler.URLShortener) server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		shortenerv1.RegisterURLShortenerServiceServer(server, shortener)
	}
}

func registerGRPCHealthService(health *handler.HealthChecker) server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		grpc_health_v1.RegisterHealthServer(server, health)
	}
}

func registerRestURLShortenerEndpoint(grpcPort string, options ...grpc.DialOption) server.RegisterEndpointFunc {
	return func(server *runtime.ServeMux) {
		err := shortenerv1.RegisterURLShortenerServiceHandlerFromEndpoint(context.Background(), server, fmt.Sprintf(":%s", grpcPort), options)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
