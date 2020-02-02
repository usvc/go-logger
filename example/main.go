package main

import (
	"gitlab.com/usvc/modules/go/logger"
)

var stdout logger.Logger
var levelledText logger.Logger
var levelledJSON logger.Logger

func init() {
	stdout = logger.New(logger.Config{
		Type: logger.TypeStdout,
	})
	levelledJSON = logger.New(logger.Config{
		Type: logger.TypeLevelled,
	})
	levelledText = logger.New(logger.Config{
		Type: logger.TypeText,
	})
}

func main() {
	stdout.Debug("hi from stdout")
	levelledJSON.Debug("hi from json")
	levelledText.Debug("hi from text")
}
