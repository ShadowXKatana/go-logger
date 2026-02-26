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

## CI/CD

### Continuous Integration (CI)

The CI workflow (`.github/workflows/ci.yml`) runs automatically on every push and pull request to any branch. It:

1. Sets up Go using the version in `go.mod`.
2. Runs all tests with coverage: `go test ./... -coverprofile=coverage.out -covermode=atomic`
3. Enforces a **70% code coverage** minimum — the workflow fails if coverage drops below this threshold.

To run the same checks locally before pushing:

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic
go tool cover -func=coverage.out
```

### Continuous Deployment (CD)

The CD workflow (`.github/workflows/cd.yml`) triggers when a tag matching `v*` is pushed. It runs the tests and then automatically creates a **GitHub Release** with generated release notes.

To cut a new release:

```bash
git tag v1.0.0
git push origin v1.0.0
```

> **Note:** The CD workflow uses the `GITHUB_PAT` repository secret. Ensure this secret is configured in **Settings → Secrets and variables → Actions** before pushing a release tag.

## License

MIT