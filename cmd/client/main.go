package main

import (
	"context"
	"fmt"
	"io"
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

	shortener := shortenerv1.NewURLShortenerServiceClient(conn)
	fmt.Println("start send...")
	send(shortener)
	fmt.Printf("end send\n\n")
	fmt.Println("start get all...")
	getAll(shortener)
	fmt.Printf("end get all\n\n")
	fmt.Println("start get detail...")
	getDetail(shortener)
	fmt.Printf("end get detail\n\n")
}

func send(shortener shortenerv1.URLShortenerServiceClient) {
	urls := []string{
		"http://this-is-a-very-long-url-1.url",
		"http://this-is-a-very-long-url-2.url",
		"http://this-is-a-very-long-url-3.url",
	}

	for _, url := range urls {
		req := &shortenerv1.CreateShortURLRequest{OriginalUrl: url}
		resp, err := shortener.CreateShortURL(context.Background(), req)
		if err != nil {
			log.Printf("create short url: %v", err)
			return
		}
		fmt.Printf("short url: %s\nexpired at: %s\n", resp.GetShortUrl(), resp.GetExpiredAt().AsTime().String())
	}
}

func getAll(shortener shortenerv1.URLShortenerServiceClient) {
	stream, err := shortener.GetAllURL(context.Background(), &shortenerv1.GetAllURLRequest{})
	if err != nil {
		log.Printf("get all: %v\n", err)
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
		fmt.Printf("short url: %s\noriginal url: %s\nexpired at: %s\n", resp.GetShortUrl(), resp.GetOriginalUrl(), resp.GetExpiredAt().AsTime().String())
	}
}

func getDetail(shortener shortenerv1.URLShortenerServiceClient) {
	stream, err := shortener.GetAllURL(context.Background(), &shortenerv1.GetAllURLRequest{})
	if err != nil {
		log.Printf("get first detail: %v\n", err)
		return
	}

	for {
		resp, serr := stream.Recv()
		if serr == io.EOF {
			break
		}
		if serr != nil {
			fmt.Printf("in get detail - stream err: %v\n", serr)
			break
		}
		url, derr := shortener.GetURLDetail(context.Background(), &shortenerv1.GetURLDetailRequest{ShortUrl: resp.GetShortUrl()})
		if derr != nil {
			log.Printf("get detail: %v\n", derr)
			return
		}
		fmt.Printf("short url: %s\noriginal url: %s\nexpired at: %s\n", url.GetShortUrl(), url.GetOriginalUrl(), url.GetExpiredAt().AsTime().String())
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
