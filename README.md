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

The CD workflow (`.github/workflows/cd.yml`) triggers when a tag matching `v*` is pushed. It:

1. Checks out the repository using the `GITHUB_PAT` secret.
2. Sets up Go using the version in `go.mod`.
3. Runs all tests to confirm the release is stable.
4. Creates a **GitHub Release** with auto-generated release notes.

#### Prerequisites — set up `GITHUB_PAT`

The workflow requires a Personal Access Token (PAT) stored as a repository secret named `GITHUB_PAT`.

1. Go to **GitHub → Settings → Developer settings → Personal access tokens → Fine-grained tokens** and click **Generate new token**.
2. Under **Repository access**, select this repository.
3. Under **Permissions → Repository permissions**, grant **Contents: Read and write**.
4. Generate the token and copy it.
5. In this repository go to **Settings → Secrets and variables → Actions → New repository secret**.
6. Set the name to `GITHUB_PAT` and paste the token as the value, then click **Add secret**.

#### Cutting a release

Once `GITHUB_PAT` is configured, push a version tag to trigger the CD workflow:

```bash
# create an annotated tag
git tag -a v1.0.0 -m "Release v1.0.0"

# push the tag to GitHub — this triggers the CD workflow
git push origin v1.0.0
```

The workflow will run, pass the tests, and publish a new release on the repository's **Releases** page.

## License

MIT