package storage

import (
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"
	"time"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	storage := &MongoStorage{}

	return fx.Module(
		"storage",
		fx.Provide(NewStorageConfig),
		fx.Provide(storage.New),
		fx.Invoke(
			func(lc fx.Lifecycle, cfg models.StorageConfig) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {

						return storage.Run(ctx, cfg)
					},
					OnStop: func(ctx context.Context) error {

						ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
						defer cancel()

						return storage.Shutdown(ctx)
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("storage")
		}),
	)
}

func NewStorageConfig() (models.StorageConfig, error) {
	var cfg models.StorageConfig
	const op = "confg storage loader failed"

	loader, err := config.NewYAML(config.File("config.yaml"))
	if err != nil {
		return cfg, fmt.Errorf("%s: %w", op, err)
	}

	err = loader.Get("storage").Populate(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("%s: %w", op, err)
	}

	return cfg, nil
}
