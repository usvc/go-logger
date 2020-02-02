package logger

import (
	"github.com/sirupsen/logrus"
)

// Level is a type that represents a logical logging level
type Level string

const (
	LevelTrace Level = "TRACE"
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

var LogrusLevelMap = map[Level]logrus.Level{
	LevelTrace: logrus.TraceLevel,
	LevelDebug: logrus.DebugLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelError: logrus.ErrorLevel,
}
