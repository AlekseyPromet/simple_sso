package app

import (
	"AlekseyPromet/authorization/internal/config"
	"AlekseyPromet/authorization/internal/grpc"
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type ApiV1 struct {
	http.Server

	logger *zap.Logger
	Config models.Config
}

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			grpc.NewModule(),
		),
		fx.Provide(config.NewConfig),
		fx.Provide(zap.NewDevelopment),
		fx.Provide(NewHttpSrever),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger,
			}
		}),
		fx.Invoke(
			func(lc fx.Lifecycle, api *ApiV1) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return api.ListenAndServe()
					},
					OnStop: func(ctx context.Context) error {
						return api.Shutdown(ctx)
					},
				})
			},
		),
	)
}

func NewHttpSrever(logger *zap.Logger, cfg models.Config) *ApiV1 {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /liveness", func(w http.ResponseWriter, r *http.Request) {
		r.Context()
		fmt.Fprintf(w, "OK")
	})

	mux.HandleFunc("GET /readness", func(w http.ResponseWriter, r *http.Request) {
		r.Context()
		fmt.Fprintf(w, "OK")
	})

	return &ApiV1{
		Server: http.Server{
			Addr:              ":" + cfg.Port,
			Handler:           mux,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		},
		logger: logger,
		Config: cfg,
	}
}
