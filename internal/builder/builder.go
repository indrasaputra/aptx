package builder

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/indrasaputra/aptx/internal/config"
	"github.com/indrasaputra/aptx/internal/http2/grpc/handler"
	"github.com/indrasaputra/aptx/internal/http2/grpc/server"
	"github.com/indrasaputra/aptx/internal/repository"
	"github.com/indrasaputra/aptx/internal/repository/cache"
	"github.com/indrasaputra/aptx/internal/repository/database"
	"github.com/indrasaputra/aptx/internal/tool"
	aptxv1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
	"github.com/indrasaputra/aptx/usecase"
)

// BuildRedisClient builds a redis client.
func BuildRedisClient(cfg config.Redis) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client := redis.NewClient(opt)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// BuildPostgresConnPool builds a SQL client.
func BuildPostgresConnPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=%d pool_max_conn_lifetime=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.MaxOpenConns,
		cfg.MaxConnLifetime,
	)

	return pgxpool.Connect(context.Background(), connCfg)
}

// BuildGRPCAptxService builds AptxService handler together with all of its dependencies.
func BuildGRPCAptxService(pool *pgxpool.Pool, rds redis.Cmdable, domain string) *handler.AptxService {
	redis := cache.NewURLRedis(rds)
	postgres := database.NewURLPostgres(pool)

	repoInserter := repository.NewURLInserter(postgres, redis)
	repoGetter := repository.NewURLGetter(postgres, redis)

	gen := tool.NewShortURLGenerator(domain)

	creator := usecase.NewShortURLCreator(gen, repoInserter)
	getter := usecase.NewURLGetter(repoGetter)

	return handler.NewAptxService(creator, getter)
}

// BuildGRPCHealthService builds HealthService handler together with all of its dependencies.
func BuildGRPCHealthService(pool *pgxpool.Pool, rds redis.Cmdable) *handler.HealthService {
	redis := cache.NewURLRedis(rds)
	postgres := database.NewURLPostgres(pool)

	repo := repository.NewHealthChecker(postgres, redis)

	checker := usecase.NewHealthChecker(repo)
	return handler.NewHealthService(checker)
}

// BuildGRPCServer builds gRPC server along with all services that needs it.
// For this project, the services are APTX and Health Checker.
// It also sets the Prometheus and Zap Logger.
func BuildGRPCServer(port string, aptx *handler.AptxService, health *handler.HealthService, options ...grpc.ServerOption) (*server.GRPC, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	grpc_zap.ReplaceGrpcLoggerV2(logger)

	options = append(options, grpc_middleware.WithUnaryServerChain(
		grpc_zap.UnaryServerInterceptor(logger),
		grpc_prometheus.UnaryServerInterceptor,
	))
	grpcServer := server.NewGRPC(port, options...)
	grpcServer.RegisterServices(
		registerGRPCPrometheus(),
		registerGRPCAptxService(aptx),
		registerGRPCHealthService(health),
	)

	return grpcServer, nil
}

// BuildRestServer builds REST server along with all services that needs it.
// For this project, there is only one service: APTX.
// Health Checker service is not included because it will only run on gRPC port.
// It also sets the Prometheus endpoint in /metrics.
func BuildRestServer(restPort, grpcPort string, options ...grpc.DialOption) (*server.Rest, error) {
	restServer := server.NewRest(restPort)
	if err := restServer.EnablePrometheus(); err != nil {
		return nil, err
	}
	err := restServer.RegisterEndpoints(
		registerRestAptxServiceEndpoint(grpcPort, options...),
	)
	if err != nil {
		return nil, err
	}
	return restServer, nil
}

func registerGRPCPrometheus() server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		grpc_prometheus.EnableHandlingTimeHistogram()
		grpc_prometheus.Register(server)
	}
}

func registerGRPCAptxService(aptx *handler.AptxService) server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		aptxv1.RegisterAptxServiceServer(server, aptx)
	}
}

func registerGRPCHealthService(health *handler.HealthService) server.RegisterServiceFunc {
	return func(server *grpc.Server) {
		grpc_health_v1.RegisterHealthServer(server, health)
	}
}

func registerRestAptxServiceEndpoint(grpcPort string, options ...grpc.DialOption) server.RegisterEndpointFunc {
	return func(server *runtime.ServeMux) error {
		return aptxv1.RegisterAptxServiceHandlerFromEndpoint(context.Background(), server, fmt.Sprintf(":%s", grpcPort), options)
	}
}
