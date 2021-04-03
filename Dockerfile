FROM golang:1.16.0-stretch AS builder
WORKDIR /app
COPY . .
RUN make compile-server
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/url-shortener-server .
COPY --from=builder /bin/grpc_health_probe ./grpc_health_probe
EXPOSE 8080
CMD ["/app/url-shortener-server"]