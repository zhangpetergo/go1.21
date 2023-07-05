package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/zhangpetergo/go1.21/slog/slogtest"
	"log/slog"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ctx := context.Background()
	// 设置上下文的内容
	v := Values{
		TraceID: uuid.NewString(),
		Now:     time.Now().UTC(),
	}
	ctx = context.WithValue(ctx, key, &v)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})
	logger := slog.New(NewContextHandler(h))

	logger.InfoCtx(ctx, "hello", "count", 3)
}
