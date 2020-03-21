# Logger

[![latest release](https://badge.fury.io/gh/usvc%2Fgo-logger.svg)](https://github.com/usvc/go-logger/releases)
[![build status](https://travis-ci.org/usvc/go-logger.svg?branch=master)](https://travis-ci.org/usvc/go-logger)
[![pipeline status](https://gitlab.com/usvc/modules/go/logger/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/logger/-/commits/master)
[![test coverage](https://api.codeclimate.com/v1/badges/3f00918cb5b964476dd0/test_coverage)](https://codeclimate.com/github/usvc/go-logger/test_coverage)
[![maintainability](https://api.codeclimate.com/v1/badges/3f00918cb5b964476dd0/maintainability)](https://codeclimate.com/github/usvc/go-logger/maintainability)

A Go package to handle logging for web services and CLI tools.

> This package wraps `github.com/sirupsen/logrus` and adds sensible defaults so that logging is as easy as `logger.New()`

| | |
| --- | --- |
| Github | [https://github.com/usvc/go-logger](https://github.com/usvc/go-logger) |
| Gitlab | [https://gitlab.com/usvc/modules/go/logger](https://gitlab.com/usvc/modules/go/logger) |

- - -

- [Logger](#logger)
- [Usage](#usage)
  - [Importing](#importing)
  - [Logging to stderr](#logging-to-stderr)
  - [Logging to file system](#logging-to-file-system)
  - [Logging to a buffer or custom io.Writer](#logging-to-a-buffer-or-custom-iowriter)
  - [Logging in JSON](#logging-in-json)
  - [Logging only Debug level and above](#logging-only-debug-level-and-above)
  - [Logging with a custom field](#logging-with-a-custom-field)
  - [Logging without levels](#logging-without-levels)
- [Configuration](#configuration)
  - [logger.Options](#loggeroptions)
- [Example Application](#example-application)
- [Development Runbook](#development-runbook)
  - [Getting Started](#getting-started)
  - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
- [License](#license)

# Usage

## Importing

```go
import (
  // ...
  "github.com/usvc/go-logger"
  // ...
)
```


## Logging to `stderr`

```go
log := logger.New(logger.Options{
  Output: logger.OutputStderr,
})
```


## Logging to file system

```go
log := logger.New(logger.Options{
  Output: logger.OutputStderr,
  OutputFilePath: "./20200223.log",
})
```

## Logging to a buffer or custom io.Writer

```go
var output bytes.Buffer
log := logger.New(logger.Options{
  Output: logger.OutputCustom,
  OutputStream: &output,
})
```

## Logging in JSON

```go
log := logger.New(logger.Options{
  Format: logger.FormatJSON,
})
```

## Logging only Debug level and above

```go
log := logger.New(logger.Options{
  Level: logger.LevelDebug,
})
```

## Logging with a custom field

```go
log := logger.New(logger.Options{
  Fields: map[string]interface{}{
    "module": "main",
  },
})
```

## Logging without levels

```go
log := logger.New(logger.Options{
  Type: logger.TypeStdout,
})
```

- - -

# Configuration

## `logger.Options`

- **`Fields`** `map[string]interface{}`: Adds custom fields to the log entry.
- **`Format`** `Format`: One of `FormatJSON` or `FormatText`. Defaults to `FormatText`.
- **`Level`** `Level`: One of `LevelTrace`, `LevelDebug`, `LevelInfo`, `LevelWarn`, or `LevelError`. Defaults to `LevelTrace`.
- **`Output`** `Output`: One of `OutputCustom`, `OutputFileSystem`, `OutputStderr`, or `OutputStdout`. Defaults to `OutputStdout`.
- **`OutputFilePath`** `string`: Path to a log file, defaults to using `os.Stdout` if file cannot be created. Only applicable when `Output` is set to `OutputFileSystem`
- **`OutputStream`** `io.Writer`: Only applicable when `Output` is set to `OutputCustom`
- **`Type`** `Type`: One of `TypeLevelled` or `TypeStdout`. Defaults to `TypeLevelled`.

- - -

# Example Application

The example application can be found at [`./cmd/logger`](./cmd/logger). To try it out from this repository, run `make run`.

To build it, run `make build_production`.

- - -

# Development Runbook

## Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

## Continuous Integration (CI) Pipeline

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

# License

Code here is licensed under the [MIT license](./LICENSE) by [@zephinzer](https://gitlab.com/zephinzer).
