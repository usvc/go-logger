package logger

import (
	"os"
	"path/filepath"

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

		output := OutputStdout
		if len(config.Output) > 0 {
			output = config.Output
		}

		log := logrus.New()

		if output == OutputFileSystem {
			outputFilePath, err := filepath.Abs(config.OutputFilePath)
			if err != nil {
				log.SetOutput(os.Stdout)
			} else {
				file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
				if err != nil {
					log.SetOutput(os.Stdout)
				} else {
					log.SetOutput(file)
				}
			}
		} else if output == OutputStderr {
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(os.Stdout)
		}
		log.SetLevel(LogrusLevelMap[level])
		log.SetReportCaller(config.ReportCaller)

		if format == FormatJSON {
			log.SetFormatter(FormatJSONPreset)
		} else {
			log.SetFormatter(FormatTextPreset)
		}
		return log
	}
	return stdoutLogger
}
