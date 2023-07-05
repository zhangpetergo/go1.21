# go1.21





该项目是为了测试 go 1.21 的新增API







## min

[min,max](https://tip.golang.org/ref/spec#Min_and_max)









## clear

[clear](https://tip.golang.org/ref/spec#Clear)



将 slice 所有的元素置为零值

删除 map 的所有映射，len(m) == 0







## slices

[slices](https://pkg.go.dev/slices@master)











## slog

slog 是 go 官方推出的新的日志记录包，期望可以解决之前的 log 包存在的痛点

- 缺少日志级别
- 不支持结构化日志，只输出纯文本消息
- 没有上下文集成的日志
- 不支持日志采样，在高吞吐的应用程序中，日志采样是减少日志量的一个重要特性
- 自定义配置有限，只支持定义输出的位置和前缀









### 使用 logger

```go
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
```

Output:

```
2023/xx/xx xx:xx:48 INFO hello count=3
2023/xx/xx xx:xx:48 hello count 3
```



默认情况下的输出格式和旧的 log 包的输出格式是一样的

default 输出格式包括，时间，日志级别，消息





### 结构化日志

slog 支持两种格式的结构化日志，text 和 json



#### Text Handler

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)
}
```

Output:

```
time=2023-xx-xxTxx:xx:xx.148+08:00 level=INFO msg=hello count=3
```

密切注意，您将看到输出格式为**键=值**对。这通常也称为[logfmt](https://brandur.org/logfmt)格式。

许多现代系统可以处理**logfmt**格式的日志。例如，[DataDog](https://docs.datadoghq.com/logs/log_configuration/parsing/?tab=matchers)、[Splunk](https://dev.splunk.com/enterprise/docs/developapps/addsupport/logging/loggingbestpractices/)、[Grafana Loki](https://grafana.com/docs/loki/latest/clients/promtail/stages/logfmt/)。Logfmt 是人类可读的并且相当容易解析。



#### JSON Handler

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)
}
```

Output:

```
{"time":"2023-xx-xxTxx:xxx:xx.6300355+08:00","level":"INFO","msg":"hello","count":3}
```

每个日志记录都是一个 JSON 对象







### 输出日志的文件和行号

```go
package main

import (
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{
		AddSource: true,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	logger.Info("hello", "count", 3)
}
```

Output:

```
time=2023-xx-xxTxx:xx:44.028+08:00 level=INFO source=slog/addsource/main.go:15 msg=hello count=3
```







### 属性

```go
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
```

Output:

```go
{"level":"INFO","msg":"Usage Statistics","current-memory":50,"min-memory":20,"max-memory":80,"cpu":10,"app-version":"v0.0.1-beta"}
```



#### 分组属性

```go
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
```

Output:

```
level=INFO msg="Usage Statistics" memory.current=50 memory.min=20 memory.max=80 cpu=10 app-version=v0.0.1-beta
```



#### 通用属性

我们希望一些字段包含在所有生成的日志中，这种属性的示例为应用名称，应用版本





