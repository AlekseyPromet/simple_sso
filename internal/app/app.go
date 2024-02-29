package app

import (
	"AlekseyPromet/authorization/internal/models"

	"go.uber.org/fx"
)

type App struct {
	port string
}

func NewApp(cfg *models.Config) *App {
	return &App{
		port: cfg.Port,
	}
}

func (a *App) Run(fx.Lifecycle) error {

	return nil
}
