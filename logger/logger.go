package logger

import (
	"log/slog"
	"os"

	"github.com/AhmedZeyad/Akalni/config"
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
