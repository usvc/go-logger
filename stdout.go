package logger

import (
	"bufio"
	"fmt"
)

type Stdout struct {
	Level        Level
	OutputStream *bufio.Writer
}

func (stdout *Stdout) AssignDefaults() {
	if stdout.OutputStream == nil {
		stdout.OutputStream = bufio.NewWriter(*DefaultOutputStream)
	}
}

func (stdout Stdout) print(args ...interface{}) {
	fmt.Fprintln(stdout.OutputStream, args...)
	stdout.OutputStream.Flush()
}

func (stdout Stdout) printf(format string, args ...interface{}) {
	fmt.Fprintf(stdout.OutputStream, format+"\n", args...)
	stdout.OutputStream.Flush()
}

func (stdout Stdout) Trace(args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelTrace] {
		stdout.print(args...)
	}
}

func (stdout Stdout) Tracef(format string, args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelTrace] {
		stdout.printf(format, args...)
	}
}

func (stdout Stdout) Debug(args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelDebug] {
		stdout.print(args...)
	}
}

func (stdout Stdout) Debugf(format string, args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelDebug] {
		stdout.printf(format, args...)
	}
}

func (stdout Stdout) Info(args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelInfo] {
		stdout.print(args...)
	}
}

func (stdout Stdout) Infof(format string, args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelInfo] {
		stdout.printf(format, args...)
	}
}

func (stdout Stdout) Warn(args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelWarn] {
		stdout.print(args...)
	}
}

func (stdout Stdout) Warnf(format string, args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelWarn] {
		stdout.printf(format, args...)
	}
}

func (stdout Stdout) Error(args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelError] {
		stdout.print(args...)
	}
}

func (stdout Stdout) Errorf(format string, args ...interface{}) {
	if LevelValueMap[stdout.Level] <= LevelValueMap[LevelError] {
		stdout.printf(format, args...)
	}
}
