.PHONY: format
format:
	bin/format.sh

.PHONY: gengrpc
gengrpc:
	bin/generate-grpc.sh

.PHONY: lint
lint:
	buf lint
	golangci-lint run ./...

.PHONY: pretty
pretty: tidy format lint

.PHONY: mockgen
mockgen:
	bin/generate-mock.sh

.PHONY: test
test:
	go test -v -race ./...

.PHONY: dep-download
dep-download:
	env GO111MODULE=on go mod download

.PHONY: tidy
tidy:
	env GO111MODULE=on go mod tidy

.PHONY: vendor
vendor:
	env GO111MODULE=on go mod vendor

.PHONY: cover
cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out 

.PHONY: coverhtml
coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: compile-server
compile-server:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o url-shortener-server cmd/server/main.go

.PHONY: docker-build-server
docker-build-server:
	docker build --no-cache -t url-shortener-server:latest .

.PHONY: docker-build-envoy
docker-build-envoy:
	docker build --no-cache -t url-shortener-envoy:latest -f bin/envoy/Dockerfile .

.PHONY: docker-build-all
docker-build-all: docker-build-server docker-build-envoy

.PHONY: docker-run-server
docker-run-server:
	docker run -p 8080:8080 -p 8081:8081 --env-file .env url-shortener-server:latest

.PHONY: docker-run-envoy
docker-run-envoy:
	docker run -p 9901:9901 -p 9090:9090 url-shortener-envoy:latest

.PHONY: docker-run-all
docker-run-all:
	docker-compose --env-file .env up

.PHONY: docker-down
docker-down:
	docker-compose --env-file .env down

.PHONY: run-prometheus
run-prometheus:
	docker run -p 9090:9090 -v ${PWD}/infrastructure/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v2.26.0

.PHONY: migration
migration:
	migrate create -ext sql -dir db/migrations $(name)

.PHONY: migrate
migrate:
	migrate -path db/migrations -database $(url) up

.PHONY: rollback
rollback:
	migrate -path db/migrations -database $(url) down 1

.PHONY: force-migrate
force-migrate:
	migrate -path db/migrations -database $(url) force $(version)