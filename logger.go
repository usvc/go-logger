package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Trace(...interface{})
	Tracef(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}

func New(config Config) Logger {
	loggerType := TypeLevelled
	if len(config.Type) > 0 {
		loggerType = config.Type
	}
	if loggerType == TypeLevelled {
		level := LevelTrace
		if len(config.Level) > 0 {
			level = config.Level
		}

		format := FormatText
		if len(config.Format) > 0 {
			format = config.Format
		}
		log := logrus.New()
		log.SetLevel(LogrusLevelMap[level])
		log.SetReportCaller(config.ReportCaller)
		if format == FormatJSON {
			log.SetFormatter(&logrus.JSONFormatter{
				CallerPrettyfier: LogrusCallerPrettyfier,
				DataKey:          FieldData,
				DisableTimestamp: false,
				FieldMap:         LogrusFieldMap,
				TimestampFormat:  TimestampFormat,
			})
		} else {
			log.SetFormatter(&logrus.TextFormatter{
				CallerPrettyfier:       LogrusCallerPrettyfier,
				DisableLevelTruncation: true,
				DisableSorting:         true,
				DisableTimestamp:       true,
				FieldMap:               LogrusFieldMap,
				FullTimestamp:          true,
				QuoteEmptyFields:       true,
				TimestampFormat:        TimestampFormat,
			})
		}
		return log
	}
	return stdoutLogger
}
