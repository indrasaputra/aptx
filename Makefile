format:
	bin/format.sh

generate:
	bin/generate.sh

lint:
	buf lint
	golangci-lint run ./...
	
pretty: tidy format lint

mockgen:
	bin/generate-mock.sh

test:
	go test -v -race ./...

dep-download:
	env GO111MODULE=on go mod download

tidy:
	env GO111MODULE=on go mod tidy

vendor:
	env GO111MODULE=on go mod vendor

cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out 

coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

compile-server:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o url-shortener cmd/server/main.go

docker-build-server:
	docker build -t url-shortener:latest .

docker-build-envoy:
	docker build -t url-shortener-envoy:latest -f bin/envoy/Dockerfile .

docker-build-all: docker-build-server docker-build-envoy

docker-run-server:
	docker run -p 8080:8080 --env-file .env url-shortener:latest

docker-run-envoy:
	docker run -p 9901:9901 -p 9090:9090 url-shortener-envoy:latest