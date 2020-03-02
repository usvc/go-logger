package logger

import (
	"bufio"
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

func New(opt ...Options) Logger {
	options := Options{}
	if len(opt) > 0 {
		options = opt[0]
	}
	options.AssignDefaults()

	level := options.Level
	format := options.Format
	output := options.Output

	switch options.Type {
	case TypeLevelled:
		log := logrus.New()
		switch output {
		case OutputFileSystem:
			outputFilePath, err := filepath.Abs(options.OutputFilePath)
			if err != nil {
				log.SetOutput(os.Stdout)
			} else if file, err := os.OpenFile(
				outputFilePath,
				OutputFileSystemFlags,
				OutputFileSystemMode,
			); err != nil {
				log.SetOutput(os.Stdout)
			} else {
				log.SetOutput(file)
			}
		case OutputStderr:
			log.SetOutput(os.Stderr)
		default:
			log.SetOutput(os.Stdout)
		}
		log.SetLevel(LogrusLevelMap[level])
		log.SetReportCaller(options.ReportCaller)

		if format == FormatJSON {
			log.SetFormatter(FormatJSONPreset)
		} else {
			log.SetFormatter(FormatTextPreset)
		}
		if options.Fields == nil {
			return log
		}
		return log.WithFields(options.Fields)
	case TypeStdout:
		fallthrough
	default:
		if output == OutputFileSystem {
			outputFilePath, err := filepath.Abs(options.OutputFilePath)
			if err == nil {
				if file, err := os.OpenFile(
					outputFilePath,
					OutputFileSystemFlags,
					OutputFileSystemMode,
				); err == nil {
					outputStream = bufio.NewWriter(file)
				}
			}
		}
		return stdoutLogger
	}
}
