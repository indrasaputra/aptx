package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/indrasaputra/url-shortener/internal/builder"
	"github.com/indrasaputra/url-shortener/internal/config"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	server := grpc.NewServer()

	urlCreator := builder.BuildGRPCURLShortener(cfg.Domain)
	shortenerv1.RegisterURLShortenerServiceServer(server, urlCreator)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	checkError(err)

	server.Serve(lis)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
