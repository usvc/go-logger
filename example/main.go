package main

import (
	"path/filepath"

	"gitlab.com/usvc/modules/go/logger"
)

var file logger.Logger
var stderr logger.Logger
var stdout logger.Logger
var levelledFileJSON logger.Logger
var levelledFileText logger.Logger
var levelledText logger.Logger
var levelledJSON logger.Logger

func init() {
	logPath, err := filepath.Abs("./example/log")
	if err != nil {
		panic(err)
	}
	file = logger.New(logger.Config{
		Output:         logger.OutputFileSystem,
		OutputFilePath: logPath,
	})
	stdout = logger.New(logger.Config{
		Type: logger.TypeStdout,
	})
	stderr = logger.New(logger.Config{
		Output: logger.OutputStderr,
		Type:   logger.TypeStdout,
	})
	levelledFileText = logger.New(logger.Config{
		Output:         logger.OutputFileSystem,
		OutputFilePath: logPath,
		Type:           logger.TypeLevelled,
		Format:         logger.FormatText,
	})
	levelledFileJSON = logger.New(logger.Config{
		Output:         logger.OutputFileSystem,
		OutputFilePath: logPath,
		Type:           logger.TypeLevelled,
		Format:         logger.FormatJSON,
	})
	levelledText = logger.New(logger.Config{
		Type: logger.TypeLevelled,
	})
	levelledJSON = logger.New(logger.Config{
		Format: logger.FormatJSON,
		Type:   logger.TypeLevelled,
	})
}

func main() {
	file.Debug("hi from file")
	stdout.Debug("hi from stdout")
	stderr.Debug("hi from stderr")
	levelledFileJSON.Debug("hi from levelled file json")
	levelledFileText.Debug("hi from levelled file text")
	levelledText.Debug("hi from levelled text")
	levelledJSON.Debug("hi from levelled json")
}
