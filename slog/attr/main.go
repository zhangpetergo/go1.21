package main

import (
	"github.com/zhangpetergo/go1.21/slog/slogtest"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))

	logger.Info("Usage Statistics",
		slog.Int("current-memory", 50),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
}
