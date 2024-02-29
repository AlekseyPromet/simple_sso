package app

import (
	"AlekseyPromet/authorization/internal/authGrpc"
	"AlekseyPromet/authorization/internal/config"
	"AlekseyPromet/authorization/internal/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type ApiV1 struct {
	*http.Server

	logger *zap.Logger
	Config models.Config
}

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			authGrpc.NewModule(),
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
						go api.Server.ListenAndServe()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
						defer cancel()

						return api.Shutdown(ctx)
					},
				})
			},
		),
	)
}

func NewHttpSrever(logger *zap.Logger, cfg models.Config) *ApiV1 {

	middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Sugar().Infof("handler %v %v", r.Method, r.URL.String())
	})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /livenness", func(w http.ResponseWriter, r *http.Request) {
		r.Context()

		middleware(w, r)
		fmt.Fprintf(w, "OK")
	})

	mux.HandleFunc("GET /readnness", func(w http.ResponseWriter, r *http.Request) {
		r.Context()

		middleware(w, r)
		fmt.Fprintf(w, "OK")
	})

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           mux,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	return &ApiV1{
		Server: srv,
		logger: logger,
		Config: cfg,
	}
}
