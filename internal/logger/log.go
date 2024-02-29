package logger

import (
	"AlekseyPromet/authorization/internal/models"
	"log/slog"
	"os"

	"go.uber.org/fx/fxevent"
)

type Logger struct {
	slogger *slog.Logger
}

var _ fxevent.Logger = &Logger{}

func NewLogger(env models.TypeEnv) *Logger {

	logger := &Logger{}

	switch env {
	case models.LocalEnv:
		logger.slogger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)

	case models.DevelopEnv:
		logger.slogger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)

	default:
		logger.slogger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelError},
			),
		)
	}

	return logger
}

func (l *Logger) LogEvent(event fxevent.Event) {
}

func (l *Logger) Shugar() *slog.Logger {
	return l.slogger
}
