FROM golang:1.16.0-stretch AS builder
WORKDIR /app
COPY . .
RUN make compile-server

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/url-shortener .
EXPOSE 8080
CMD ["/app/url-shortener"]