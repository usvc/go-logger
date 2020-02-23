package logger

import (
	"github.com/sirupsen/logrus"
)

// Level is a type that represents a logical log level
type Level string

// Levels is a slice of levels
type Levels []Level

// Includes returns true if the provided :this is a proper log level
func (levels Levels) Includes(this Level) bool {
	for _, l := range levels {
		if l == this {
			return true
		}
	}
	return false
}

const (
	LevelTrace Level = "trace"
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"

	DefaultLevel = LevelTrace
)

var ValidLevels = Levels{
	LevelTrace,
	LevelDebug,
	LevelInfo,
	LevelWarn,
	LevelError,
}

var LogrusLevelMap = map[Level]logrus.Level{
	LevelTrace: logrus.TraceLevel,
	LevelDebug: logrus.DebugLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelError: logrus.ErrorLevel,
}
