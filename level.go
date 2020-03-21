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
	// LevelTrace defines a level for logs that only developers should see
	LevelTrace Level = "trace"
	// LevelDebug defines a level for logs that should appear in production
	// logs but are not indicative of business flow success/failure
	LevelDebug Level = "debug"
	// LevelInfo defines a level for logs that are indicative of business
	// flow success/failure
	LevelInfo Level = "info"
	// LevelWarn defines a level of logs for events that may disrupt
	// flow and indicate that the code didn't execute perfectly as expected
	LevelWarn Level = "warn"
	// LevelError defines a level of logs for events that are disruptive
	// to the application and which may result in the application exitting
	LevelError Level = "error"

	DefaultLevel = LevelTrace
)

// ValidLevels defines the levels available for use
var ValidLevels = Levels{
	LevelTrace,
	LevelDebug,
	LevelInfo,
	LevelWarn,
	LevelError,
}

// LevelvalueMap maps levels to numbers for loggers to determine
// whether a log should be displayed based on it's level
var LevelValueMap = map[Level]uint{
	LevelTrace: 100,
	LevelDebug: 200,
	LevelInfo:  300,
	LevelWarn:  400,
	LevelError: 500,
}

// LogrusLevelMap maps levels according to this package, to the
// github.com/sirupsen/logrus package's levels
var LogrusLevelMap = map[Level]logrus.Level{
	LevelTrace: logrus.TraceLevel,
	LevelDebug: logrus.DebugLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelError: logrus.ErrorLevel,
}
