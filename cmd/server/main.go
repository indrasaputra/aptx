package main

import (
	"google.golang.org/grpc"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
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

	grpcServer, gerr := builder.BuildGRPCServer(cfg.PortGRPC, shortener, health)
	checkError(gerr)
	_ = grpcServer.Run()

	restServer, herr := builder.BuildRestServer(cfg.PortHTTP, cfg.PortGRPC, grpc.WithInsecure())
	checkError(herr)
	_ = restServer.Run()

	_ = grpcServer.AwaitTermination()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
