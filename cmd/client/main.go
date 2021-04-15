package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/indrasaputra/aptx/internal/config"
	aptxv1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	conn, err := grpc.Dial(fmt.Sprintf(":%s", cfg.PortGRPC), grpc.WithInsecure())
	checkError(err)
	defer conn.Close()

	aptx := aptxv1.NewURLShortenerServiceClient(conn)
	fmt.Println("start send...")
	send(aptx)
	fmt.Printf("end send\n\n")
	fmt.Println("start get all...")
	streamAll(aptx)
	fmt.Printf("end get all\n\n")
	fmt.Println("start get detail...")
	getDetail(aptx)
	fmt.Printf("end get detail\n\n")
}

func send(aptx aptxv1.URLShortenerServiceClient) {
	urls := []string{
		"http://this-is-a-very-long-url-1.url",
		"http://this-is-a-very-long-url-2.url",
		"http://this-is-a-very-long-url-3.url",
	}

	for _, url := range urls {
		req := &aptxv1.CreateShortURLRequest{OriginalUrl: url}
		resp, err := aptx.CreateShortURL(context.Background(), req)
		if err != nil {
			log.Printf("create short url: %v", err)
			return
		}
		fmt.Printf("short url: %s\nexpired at: %s\n", resp.GetUrl().GetShortUrl(), resp.GetUrl().GetExpiredAt().AsTime().String())
	}
}

func streamAll(aptx aptxv1.URLShortenerServiceClient) {
	stream, err := aptx.StreamAllURL(context.Background(), &aptxv1.StreamAllURLRequest{})
	if err != nil {
		log.Printf("stream all: %v\n", err)
		return
	}

	for {
		resp, serr := stream.Recv()
		if serr == io.EOF {
			break
		}
		if serr != nil {
			fmt.Printf("stream err: %v\n", serr)
			break
		}
		fmt.Printf("short url: %s\noriginal url: %s\nexpired at: %s\n", resp.GetUrl().GetShortUrl(), resp.GetUrl().GetOriginalUrl(), resp.GetUrl().GetExpiredAt().AsTime().String())
	}
}

func getDetail(aptx aptxv1.URLShortenerServiceClient) {
	resp, err := aptx.GetAllURL(context.Background(), &aptxv1.GetAllURLRequest{})
	if err != nil {
		log.Printf("get first detail: %v\n", err)
		return
	}

	for _, res := range resp.GetUrls() {
		url, derr := aptx.GetURLDetail(context.Background(), &aptxv1.GetURLDetailRequest{Code: res.GetCode()})
		if derr != nil {
			log.Printf("get detail: %v\n", derr)
			return
		}
		fmt.Printf("short url: %s\noriginal url: %s\nexpired at: %s\n", url.GetUrl().GetShortUrl(), url.GetUrl().GetOriginalUrl(), url.GetUrl().GetExpiredAt().AsTime().String())
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
