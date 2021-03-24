package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/indrasaputra/url-shortener/internal/config"
	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	conn, err := grpc.Dial(fmt.Sprintf(":%s", cfg.Port), grpc.WithInsecure())
	checkError(err)
	defer conn.Close()

	urlCreator := shortenerv1.NewURLShortenerServiceClient(conn)

	req := &shortenerv1.CreateShortURLRequest{OriginalUrl: "http://this-is-a-very-long-url.url"}
	resp, err := urlCreator.CreateShortURL(context.Background(), req)
	if err != nil {
		log.Printf("create short url: %s", err.Error())
		return
	}
	fmt.Printf("short url: %s\nexpired at: %s\n", resp.GetShortUrl(), resp.GetExpiredAt().AsTime().String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
