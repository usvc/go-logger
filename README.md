# Logger

[![latest release](https://badge.fury.io/gh/usvc%2Fgo-logger.svg)](https://github.com/usvc/go-logger/releases)
[![pipeline status](https://gitlab.com/usvc/modules/go/logger/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/logger/-/commits/master)
[![release status](https://travis-ci.org/usvc/go-logger.svg?branch=master)](https://travis-ci.org/usvc/go-logger)

A Go package to handle logging for web services and CLI tools.

| | |
| --- | --- |
| Github | [https://github.com/usvc/go-logger](https://github.com/usvc/go-logger) |
| Gitlab | [https://gitlab.com/usvc/modules/go/logger](https://gitlab.com/usvc/modules/go/logger) |


- [Logger](#logger)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Logging to stdout](#logging-to-stdout)
    - [Logging to stderr](#logging-to-stderr)
    - [Logging to file system](#logging-to-file-system)
    - [Logging in JSON](#logging-in-json)
    - [Logging without levels](#logging-without-levels)
    - [Logging without caller functions](#logging-without-caller-functions)
  - [Example Application](#example-application)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
  - [License](#license)

## Usage

### Importing

```go
import (
  // ...
  "github.com/usvc/go-logger"
  // ...
)
```

### Logging to `stdout`

```go
log := logger.New(logger.Options{
  Type: logger.TypeStdout,
})
```


### Logging to `stderr`

```go
log := logger.New(logger.Options{
  Output: logger.OutputStderr,
})
```


### Logging to file system

```go
log := logger.New(logger.Options{
  Output: logger.OutputStderr,
  OutputFilePath: "./20200223.log",
})
```

### Logging in JSON

```go
log := logger.New(logger.Options{
  Format: logger.FormatJSON,
})
```

### Logging without levels

```go
log := logger.New(logger.Options{
  Type: logger.TypeStdout,
})
```

### Logging without caller functions

```go
log := logger.New(logger.Options{
  ReportCaller: true,
})
```

- - -

## Example Application

The example application can be found at [`./cmd/logger`](./cmd/logger). To try it out from this repository, run `make run`.

To build it, run `make build_production`.

- - -

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
