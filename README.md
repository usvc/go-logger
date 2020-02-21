# Logger

[![pipeline status](https://gitlab.com/usvc/modules/go/logger/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/logger/-/commits/master)

Generic all-purpose logger.

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

## License

Code here is licensed under the [MIT license](./LICENSE) by [@zephinzer](https://gitlab.com/zephinzer).
