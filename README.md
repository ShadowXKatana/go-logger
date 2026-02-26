# go-logger

A simple leveled logger for Go that emits structured JSON logs, backed by [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap).

## Installation

```bash
go get github.com/ShadowXKatana/go-logger
```

## Usage

```go
package main

import logger "github.com/ShadowXKatana/go-logger"

func main() {
    log := logger.New(logger.INFO)

    log.Debug("this won't appear")
    log.Info("application started")
    log.Warn("low memory")
    log.Error("something went wrong")

    log.Debugf("connecting to %s", "localhost")
    log.Infof("user %s logged in", "alice")
    log.Warnf("retrying %d more times", 3)
    log.Errorf("failed to open %s", "config.yaml")
}
```

## Log Levels

| Level | Value |
|-------|-------|
| DEBUG | 0     |
| INFO  | 1     |
| WARN  | 2     |
| ERROR | 3     |

Only messages at or above the configured level are emitted.

## API

### Constructor

| Function | Description |
|----------|-------------|
| `New(level Level) *Logger` | Create a new logger that writes JSON to stdout at the given minimum level. |

### Methods

| Method | Description |
|--------|-------------|
| `Debug(msg string)` | Log a message at DEBUG level. |
| `Info(msg string)` | Log a message at INFO level. |
| `Warn(msg string)` | Log a message at WARN level. |
| `Error(msg string)` | Log a message at ERROR level. |
| `Debugf(format string, args ...any)` | Log a formatted message at DEBUG level. |
| `Infof(format string, args ...any)` | Log a formatted message at INFO level. |
| `Warnf(format string, args ...any)` | Log a formatted message at WARN level. |
| `Errorf(format string, args ...any)` | Log a formatted message at ERROR level. |

## License

MIT