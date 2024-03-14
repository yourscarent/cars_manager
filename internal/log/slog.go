package log

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func MustSetup(env string) Logger {
	var l *slog.Logger

	switch env {
	case "local":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		outputFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			panic("failed to open logs file: " + err.Error())
		}
		l = slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, outputFile), &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return Logger{l}
}
