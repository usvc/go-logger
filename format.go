package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Format string

const (
	FormatText Format = "TEXT"
	FormatJSON Format = "JSON"
)

var FormatJSONPreset = &logrus.JSONFormatter{
	CallerPrettyfier: LogrusCallerPrettyfier,
	DataKey:          FieldData,
	DisableTimestamp: false,
	FieldMap:         LogrusFieldMap,
	TimestampFormat:  TimestampFormat,
}

var FormatTextPreset = &logrus.TextFormatter{
	CallerPrettyfier:       LogrusCallerPrettyfier,
	DisableLevelTruncation: true,
	DisableSorting:         true,
	DisableTimestamp:       false,
	FieldMap:               LogrusFieldMap,
	FullTimestamp:          true,
	QuoteEmptyFields:       true,
	TimestampFormat:        TimestampFormat,
}

func LogrusCallerPrettyfier(frame *runtime.Frame) (function string, file string) {
	return frame.Function, filepath.Base(frame.File)
}
