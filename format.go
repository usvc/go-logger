package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Format defines a valid log format
type Format string

// Formats is a slice of Format
type Formats []Format

// Includes returns true if the provided :this is a proper log format
func (formats Formats) Includes(this Format) bool {
	for _, f := range formats {
		if f == this {
			return true
		}
	}
	return false
}

const (
	FormatText Format = "text"
	FormatJSON Format = "json"

	DefaultFormat = FormatText
)

var ValidFormats Formats = Formats{FormatJSON, FormatText}

var FormatJSONPreset = &logrus.JSONFormatter{
	CallerPrettyfier: LogrusCallerPrettyfier,
	DataKey:          FieldData,
	DisableTimestamp: false,
	FieldMap:         LogrusFieldMap,
	TimestampFormat:  TimestampFormat,
}

var FormatTextPreset = &logrus.TextFormatter{
	CallerPrettyfier:       LogrusCallerPrettyfier,
	DisableLevelTruncation: false,
	DisableSorting:         true,
	DisableTimestamp:       false,
	FieldMap:               LogrusFieldMap,
	FullTimestamp:          true,
	QuoteEmptyFields:       true,
	TimestampFormat:        TimestampFormat,
}

func LogrusCallerPrettyfier(frame *runtime.Frame) (function string, file string) {
	return filepath.Base(frame.Function), filepath.Base(frame.File)
}
