package grpc

import (
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"grpc",
		fx.Provide(
			NewGrpcConfig,
		),
		fx.Invoke(
			func(lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("grpc")
		}),
	)
}

func NewGrpcConfig(provider config.Provider) (models.GrpcConfig, error) {
	var cfg models.GrpcConfig
	err := provider.Get("grpc").Populate(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("grpc config: %w", err)
	}

	return cfg, nil
}
