package authGrpc

import (
	"AlekseyPromet/authorization/internal/models"
	"AlekseyPromet/authorization/internal/storage"
	"context"
	"fmt"
	"time"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewModule() fx.Option {

	service := &GrpcServerApiV1{}

	return fx.Module(
		"grpc server",
		fx.Options(
			storage.NewModule(),
		),
		fx.Provide(NewGrpcConfig),
		fx.Invoke(
			func(lc fx.Lifecycle, cfg models.GrpcConfig, storage *storage.MongoStorage) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {

						Register(grpc.NewServer())

						go service.Run(cfg, storage)

						return nil
					},
					OnStop: func(ctx context.Context) error {

						ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
						defer cancel()

						return service.Shutdown(ctx)
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("grpc")
		}),
	)
}

func NewGrpcConfig() (models.GrpcConfig, error) {
	var cfg models.GrpcConfig
	const op = "confg grpc loader failed"

	loader, err := config.NewYAML(config.File("config.yaml"))
	if err != nil {
		return cfg, fmt.Errorf("%s: %w", op, err)
	}

	err = loader.Get("grpc").Populate(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("%s: %w", op, err)
	}

	return cfg, nil
}
