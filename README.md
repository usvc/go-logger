# Logger

## `stdout` logger

```go
log := logger.New(logger.Config{
  Type: logger.TypeStdout,
})
```


## `stderr` logger

```go
log := logger.New(logger.Config{
  Output: logger.OutputStderr,
  Type:   logger.TypeStdout,
})
```

## Levelled logger (Text)

```go
log := logger.New(logger.Config{
  ReportCaller: true,
  Type:         logger.TypeLevelled,
})
```

## Levelled Logger (JSON)

```go
log := logger.New(logger.Config{
  Format: logger.FormatJSON,
  Type:   logger.TypeLevelled,
})
```
