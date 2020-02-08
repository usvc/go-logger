package logger

import "fmt"

var stdoutLogger = Stdout{}

type Stdout struct{}

var stdoutPrint = stdoutPrinter
var stdoutPrintf = stdoutPrinterf

func stdoutPrinter(args ...interface{}) {
	fmt.Println(args...)
}

func stdoutPrinterf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s\n", format), args...)
}

func (stdout Stdout) Trace(args ...interface{}) {
	stdoutPrint(args...)
}

func (stdout Stdout) Tracef(format string, args ...interface{}) {
	stdoutPrintf(format, args...)
}

func (stdout Stdout) Debug(args ...interface{}) {
	stdoutPrint(args...)
}

func (stdout Stdout) Debugf(format string, args ...interface{}) {
	stdoutPrintf(format, args...)
}

func (stdout Stdout) Info(args ...interface{}) {
	stdoutPrint(args...)
}

func (stdout Stdout) Infof(format string, args ...interface{}) {
	stdoutPrintf(format, args...)
}

func (stdout Stdout) Warn(args ...interface{}) {
	stdoutPrint(args...)
}

func (stdout Stdout) Warnf(format string, args ...interface{}) {
	stdoutPrintf(format, args...)
}

func (stdout Stdout) Error(args ...interface{}) {
	stdoutPrint(args...)
}

func (stdout Stdout) Errorf(format string, args ...interface{}) {
	stdoutPrintf(format, args...)
}
