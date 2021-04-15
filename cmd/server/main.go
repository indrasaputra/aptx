package main

import (
	"google.golang.org/grpc"

	"github.com/indrasaputra/aptx/internal/builder"
	"github.com/indrasaputra/aptx/internal/config"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	postgres, perr := builder.BuildPostgresConnPool(cfg.Postgres)
	checkError(perr)
	redis, rerr := builder.BuildRedisClient(cfg.Redis)
	checkError(rerr)

	aptx := builder.BuildGRPCAptxService(postgres, redis, cfg.Domain)
	health := builder.BuildGRPCHealthService(postgres, redis)

	grpcServer, gerr := builder.BuildGRPCServer(cfg.PortGRPC, aptx, health)
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
