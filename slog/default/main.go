package main

import (
	"log"
	"log/slog"
)

func main() {

	// default 输出
	slog.Info("hello", "count", 3)
	log.Println("hello", "count", 3)
}
