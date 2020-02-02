package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	FieldData      = "@data"
	FieldTimestamp = "@timestamp"
	FieldLevel     = "@level"
	FieldMessage   = "@message"
	FieldFunc      = "@caller"

	TimestampFormat = "20060102150405"
)

var LogrusFieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  FieldTimestamp,
	logrus.FieldKeyLevel: FieldLevel,
	logrus.FieldKeyMsg:   FieldMessage,
	logrus.FieldKeyFunc:  FieldFunc,
}
