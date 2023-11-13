package logger

import "github.com/gookit/slog"

func Get() *slog.Logger {
	return loggerInstance
}
