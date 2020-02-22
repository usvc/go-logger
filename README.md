# Logger

[![pipeline status](https://gitlab.com/usvc/modules/go/logger/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/logger/-/commits/master)

A Go package to handle logging for web services and CLI tools

- [Logger](#logger)
  - [Usage](#usage)
    - [Import](#import)
    - [stdout logger](#stdout-logger)
    - [stderr logger](#stderr-logger)
    - [Levelled logger (Text)](#levelled-logger-text)
    - [Levelled Logger (JSON)](#levelled-logger-json)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
  - [License](#license)

## Usage

### Import

```go
import (
  // ...
  "github.com/usvc/logger"
  // ...
)
```

### `stdout` logger

```go
log := logger.New(logger.Config{
  Type: logger.TypeStdout,
})
```


### `stderr` logger

```go
log := logger.New(logger.Config{
  Output: logger.OutputStderr,
  Type:   logger.TypeStdout,
})
```

### Levelled logger (Text)

```go
log := logger.New(logger.Config{
  ReportCaller: true,
  Type:         logger.TypeLevelled,
})

// file based
logPath, _ := filepath.Abs("./example/log.text")
log := logger.New(logger.Config{
  Output:         logger.OutputFileSystem,
  OutputFilePath: logPath,
  ReportCaller:   true,
  Type:           logger.TypeLevelled,
})
```

### Levelled Logger (JSON)

```go
log := logger.New(logger.Config{
  Format: logger.FormatJSON,
  Type:   logger.TypeLevelled,
})

// file based
logPath, _ := filepath.Abs("./example/log.json")
log := logger.New(logger.Config{
  Format:         logger.FormatJSON,
  Output:         logger.OutputFileSystem,
  OutputFilePath: logPath,
  ReportCaller:   true,
  Type:           logger.TypeLevelled,
})
```

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

## License

Code here is licensed under the [MIT license](./LICENSE) by [@zephinzer](https://gitlab.com/zephinzer).
