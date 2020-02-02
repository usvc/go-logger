package logger

import "fmt"

var stdoutLogger = Stdout{}

type Stdout struct{}

func stdoutPrinter(args ...interface{}) {
	fmt.Print(args...)
}

func stdoutPrinterf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (stdout Stdout) Trace(args ...interface{}) {
	stdoutPrinter(args...)
}

func (stdout Stdout) Tracef(format string, args ...interface{}) {
	stdoutPrinterf(format, args...)
}

func (stdout Stdout) Debug(args ...interface{}) {
	stdoutPrinter(args...)
}

func (stdout Stdout) Debugf(format string, args ...interface{}) {
	stdoutPrinterf(format, args...)
}

func (stdout Stdout) Info(args ...interface{}) {
	stdoutPrinter(args...)
}

func (stdout Stdout) Infof(format string, args ...interface{}) {
	stdoutPrinterf(format, args...)
}

func (stdout Stdout) Warn(args ...interface{}) {
	stdoutPrinter(args...)
}

func (stdout Stdout) Warnf(format string, args ...interface{}) {
	stdoutPrinterf(format, args...)
}

func (stdout Stdout) Error(args ...interface{}) {
	stdoutPrinter(args...)
}

func (stdout Stdout) Errorf(format string, args ...interface{}) {
	stdoutPrinterf(format, args...)
}
