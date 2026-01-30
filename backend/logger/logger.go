package logger

import (
	"log/slog"
	"os"

	"github.com/ba7rIbrahim/Akalni/config"
)

var Log *slog.Logger

func Init(config *config.Config) {
	option := slog.HandlerOptions{
		AddSource: true,
	}
	if config.ISDev == "true" {

		option.Level = slog.LevelDebug
	} else {
		option.Level = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stdout, &option)
	Log = slog.New(handler)
	slog.SetDefault(Log)
}

func Error(message string, args ...any) {
	Log.Error(message, args...)
}
func Info(message string, args ...any) {
	Log.Info(message, args...)
}
func Warn(message string, args ...any) {
	Log.Warn(message, args...)
}
func Debug(message string, args ...any) {
	Log.Debug(message, args...)
}
