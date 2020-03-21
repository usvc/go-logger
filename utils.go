package logger

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func New(opt ...Options) Logger {
	options := Options{}
	if len(opt) > 0 {
		options = opt[0]
	}
	options.AssignDefaults()

	switch options.Type {
	case TypeLevelled:
		return newLevelled(options)
	case TypeStdout:
		fallthrough
	default:
		return newStdout(options)
	}
}

func newStdout(options Options) Logger {
	level := options.Level
	output := options.Output

	log := Stdout{Level: level}
	switch output {
	case OutputFileSystem:
		outputFilePath, err := filepath.Abs(options.OutputFilePath)
		if err == nil {
			if file, err := os.OpenFile(
				outputFilePath,
				OutputFileSystemFlags,
				OutputFileSystemMode,
			); err == nil {
				log.OutputStream = bufio.NewWriter(file)
			}
		}
	case OutputCustom:
		if options.OutputStream != nil {
			log.OutputStream = bufio.NewWriter(options.OutputStream)
		} else {
			log.OutputStream = bufio.NewWriter(*DefaultOutputStream)
		}
	case OutputStderr:
		log.OutputStream = bufio.NewWriter(os.Stderr)
	case OutputStdout:
		fallthrough
	default:
		log.OutputStream = bufio.NewWriter(*DefaultOutputStream)
	}
	return log
}

func newLevelled(options Options) Logger {
	level := options.Level
	format := options.Format
	output := options.Output

	log := logrus.New()

	log.SetLevel(LogrusLevelMap[level])
	log.SetReportCaller(true)

	switch output {
	case OutputFileSystem:
		var outputFilePath string
		var fileHandler *os.File
		var err error
		if outputFilePath, err = filepath.Abs(options.OutputFilePath); err != nil {
			log.SetOutput(*DefaultOutputStream)
			break
		}
		if fileHandler, err = os.OpenFile(
			outputFilePath,
			OutputFileSystemFlags,
			OutputFileSystemMode,
		); err != nil {
			log.SetOutput(*DefaultOutputStream)
			break
		}
		log.SetOutput(fileHandler)
	case OutputCustom:
		if options.OutputStream != nil {
			log.SetOutput(options.OutputStream)
		} else {
			log.SetOutput(*DefaultOutputStream)
		}
	case OutputStderr:
		log.SetOutput(os.Stderr)
	case OutputStdout:
		fallthrough
	default:
		log.SetOutput(*DefaultOutputStream)
	}

	if format == FormatJSON {
		log.SetFormatter(FormatJSONPreset)
	} else {
		log.SetFormatter(FormatTextPreset)
	}

	if options.Fields == nil {
		return log
	}
	return log.WithFields(options.Fields)
}
