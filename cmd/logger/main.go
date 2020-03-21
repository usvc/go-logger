package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/go-logger"
)

var (
	format         string
	level          string
	output         string
	outputFilePath string
	reportCaller   bool
	logType        string
)
var options = logger.Options{}
var log logger.Logger
var cmd = &cobra.Command{
	Use:  "logger",
	Long: "logger - sample application to demonstrate logging with github.com/usvc/go-logger",
	Example: strings.ReplaceAll(
		strings.Trim(`
			logger --format 'json'
			logger --level 'info'
			logger --output 'stderr'
			logger --output-path './log.txt'
			logger --report-caller
			logger --type 'stdout'`, "\n",
		), "\t", " ",
	),
	PreRun: preRun,
	Run:    run,
}

func preRun(command *cobra.Command, args []string) {
	fmt.Println("configuration")
	fmt.Println("^^^^^^^^^^^^^")
	options.Format = logger.Format(format)
	options.Level = logger.Level(level)
	options.Output = logger.Output(output)
	options.OutputFilePath = outputFilePath
	options.Type = logger.Type(logType)
	options.AssignDefaults()
	fmt.Println("{{field}}     : {{input}} > {{evaluated}}")
	fmt.Printf("format        : %s > %s\n", format, options.Format)
	fmt.Printf("level         : %s > %s\n", level, options.Level)
	fmt.Printf("output        : %s > %s\n", output, options.Output)
	fmt.Printf("output-path   : %s > %s\n", outputFilePath, options.OutputFilePath)
	fmt.Printf("type          : %s > %s\n", logType, options.Type)
	log = logger.New(options)
}

func run(command *cobra.Command, args []string) {
	fmt.Println("")
	fmt.Println("sample logs")
	fmt.Println("^^^^^^^^^^^")
	log.Trace("examle trace")
	log.Debug("example debug")
	log.Info("example info")
	log.Warn("example warn")
	log.Error("example error")
}

func init() {
	cmd.Flags().StringVarP(&format, "format", "f", string(logger.DefaultFormat), fmt.Sprintf("specifies the log format {oneof:%v}", logger.ValidFormats))
	cmd.Flags().StringVarP(&level, "level", "l", string(logger.DefaultLevel), fmt.Sprintf("specifies the log level {oneof:%v}", logger.ValidLevels))
	cmd.Flags().StringVarP(&output, "output", "o", string(logger.DefaultOutput), fmt.Sprintf("specifies the type of output {oneof:%v}", logger.ValidOutputs))
	cmd.Flags().StringVarP(&logType, "type", "t", string(logger.DefaultType), fmt.Sprintf("specifies the type of logger {oneof:%v}", logger.ValidTypes))
	cmd.Flags().StringVarP(&outputFilePath, "output-path", "O", logger.DefaultOutputFilePath, fmt.Sprintf("specifies the location of the output file path"))
}

func main() {
	cmd.Execute()
}
