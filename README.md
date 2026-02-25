# go-logger

A simple leveled logger for Go.

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

    log.Infof("user %s logged in", "alice")
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

## License

MIT