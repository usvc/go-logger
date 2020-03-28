# Logger

[![latest release](https://badge.fury.io/gh/usvc%2Fgo-logger.svg)](https://github.com/usvc/go-logger/releases)
[![build status](https://travis-ci.org/usvc/go-logger.svg?branch=master)](https://travis-ci.org/usvc/go-logger)
[![pipeline status](https://gitlab.com/usvc/modules/go/logger/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/logger/-/commits/master)
[![test coverage](https://api.codeclimate.com/v1/badges/3f00918cb5b964476dd0/test_coverage)](https://codeclimate.com/github/usvc/go-logger/test_coverage)
[![maintainability](https://api.codeclimate.com/v1/badges/3f00918cb5b964476dd0/maintainability)](https://codeclimate.com/github/usvc/go-logger/maintainability)

A Go package to handle logging for web services and CLI tools.

> This package wraps `github.com/sirupsen/logrus` and adds defaults so that creating a sensible logger is as easy as `logger.New()`

| | |
| --- | --- |
| Github | [https://github.com/usvc/go-logger](https://github.com/usvc/go-logger) |
| Gitlab | [https://gitlab.com/usvc/modules/go/logger](https://gitlab.com/usvc/modules/go/logger) |

- - -

- [Logger](#logger)
- [Usage](#usage)
  - [Importing](#importing)
  - [Instantiating a basic logger](#instantiating-a-basic-logger)
  - [Instantiating a logger without output](#instantiating-a-logger-without-output)
  - [Instantiating a basic logger instance](#instantiating-a-basic-logger-instance)
  - [Logging to stderr](#logging-to-stderr)
  - [Logging to file system](#logging-to-file-system)
  - [Logging to a buffer or custom io.Writer](#logging-to-a-buffer-or-custom-iowriter)
  - [Logging in JSON](#logging-in-json)
  - [Logging only Debug level and above](#logging-only-debug-level-and-above)
  - [Logging with a custom field](#logging-with-a-custom-field)
  - [Logging without levels](#logging-without-levels)
- [Documentation](#documentation)
  - [Configuration](#configuration)
    - [logger.Options](#loggeroptions)
  - [Constants](#constants)
    - [Format](#format)
    - [Level](#level)
    - [Output](#output)
    - [Type](#type)
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

## Instantiating a basic logger

```go
log := logger.New()
```


## Instantiating a logger without output

```go
log := logger.New(logger.Options{
  Type: logger.TypeNoOp,
})
```


## Instantiating a basic logger instance

```go
log := logger.NewLogrusEntry()
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

# Documentation

## Configuration

### `logger.Options`

- **`Fields`** `map[string]interface{}`: Adds custom fields to the log entry.
- **`Format`** `logger.Format`: One of `FormatJSON` or `FormatText`. Defaults to `FormatText`.
- **`Level`** `logger.Level`: One of `LevelTrace`, `LevelDebug`, `LevelInfo`, `LevelWarn`, or `LevelError`. Defaults to `LevelTrace`.
- **`Output`** `logger.Output`: One of `OutputCustom`, `OutputFileSystem`, `OutputStderr`, or `OutputStdout`. Defaults to `OutputStdout`.
- **`OutputFilePath`** `string`: Path to a log file, defaults to using `os.Stdout` if file cannot be created. Only applicable when `Output` is set to `OutputFileSystem`
- **`OutputStream`** `io.Writer`: Only applicable when `Output` is set to `OutputCustom`
- **`Type`** `logger.Type`: One of `TypeLevelled` or `TypeStdout`. Defaults to `TypeLevelled`.

## Constants

### `Format`

Format defines how the logs output should be formatted.

- **`FormatText`**: output plain text
- **`FormatJSON`**: output JSON-formatted text

### `Level`

Level defines the level of the logs.

- **`LevelTrace`**: all other logs
- **`LevelDebug`**: logs related to code execution
- **`LevelInfo`**: logs related to business-flow success
- **`LevelWarn`**: logs related to business-flow errors
- **`LevelError`**: logs related to system-level failures

### `Output`

Output defines where the logs should be streamed to.

- **`OutputCustom`**: send logs to a custom `io.Writer`
- **`OutputFileSystem`**: send logs to a file
- **`OutputStderr`**: send logs to standard error
- **`OutputStdout`**: send logs to standard output

### `Type`

Type defines the type of logger desired.

- **`TypeLevelled`**: defines a levelled logger
- **`TypeStdout`**: defines a plaintext logger

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
