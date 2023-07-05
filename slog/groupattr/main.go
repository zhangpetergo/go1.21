package main

import (
	"github.com/zhangpetergo/go1.21/slog/slogtest"
	"log/slog"
	"os"
)

func main() {

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))

	logger.Info("Usage Statistics",
		slog.Group("memory", slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)

}
