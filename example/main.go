package main

import (
	"gitlab.com/usvc/modules/go/logger"
)

var stderr logger.Logger
var stdout logger.Logger
var levelledText logger.Logger
var levelledJSON logger.Logger

func init() {
	stdout = logger.New(logger.Config{
		Type: logger.TypeStdout,
	})
	stderr = logger.New(logger.Config{
		Output: logger.OutputStderr,
		Type:   logger.TypeStdout,
	})
	levelledJSON = logger.New(logger.Config{
		Format: logger.FormatJSON,
		Type:   logger.TypeLevelled,
	})
	levelledText = logger.New(logger.Config{
		ReportCaller: true,
		Type:         logger.TypeLevelled,
	})
}

func main() {
	stderr.Debug("hi from stderr")
	stdout.Debug("hi from stdout")
	levelledJSON.Debug("hi from json")
	levelledText.Debug("hi from text")
	levelledJSON.Debug("hi from json")
	levelledText.Debug("hi from text")
	levelledJSON.Debug("hi from json")
	levelledText.Debug("hi from text")
	levelledJSON.Debug("hi from json")
}
