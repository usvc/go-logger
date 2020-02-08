package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StdoutTests struct {
	suite.Suite
}

func TestStdout(t *testing.T) {
	suite.Run(t, &StdoutTests{})
}

func (s *StdoutTests) Test_stdoutPrinter() {
	stdout := captureStdoutFrom(func() {
		stdoutPrinter("hello")
	})
	s.Equal("hello\n", stdout)
}

func (s *StdoutTests) Test_stdoutPrinterf() {
	stdout := captureStdoutFrom(func() {
		stdoutPrinterf("hello %s", "world")
	})
	s.Equal("hello world\n", stdout)
}

func (s *StdoutTests) TestStdout() {
	stdout := Stdout{}
	s.True(isStdoutPrintCalled(func() {
		stdout.Trace("trace")
	}))
	s.True(isStdoutPrintfCalled(func() {
		stdout.Tracef("trace")
	}))
	s.True(isStdoutPrintCalled(func() {
		stdout.Debug("debug")
	}))
	s.True(isStdoutPrintfCalled(func() {
		stdout.Debugf("debug")
	}))
	s.True(isStdoutPrintCalled(func() {
		stdout.Info("info")
	}))
	s.True(isStdoutPrintfCalled(func() {
		stdout.Infof("info")
	}))
	s.True(isStdoutPrintCalled(func() {
		stdout.Warn("warn")
	}))
	s.True(isStdoutPrintfCalled(func() {
		stdout.Warnf("warn")
	}))
	s.True(isStdoutPrintCalled(func() {
		stdout.Error("error")
	}))
	s.True(isStdoutPrintfCalled(func() {
		stdout.Errorf("error")
	}))
}

func isStdoutPrintCalled(whenThisIsCalled func()) bool {
	originalStdoutPrint := stdoutPrint
	defer func() {
		stdoutPrint = originalStdoutPrint
	}()
	stdoutPrintCallCount := 0
	stdoutPrint = func(...interface{}) {
		stdoutPrintCallCount++
	}
	whenThisIsCalled()
	return stdoutPrintCallCount == 1
}

func isStdoutPrintfCalled(whenThisIsCalled func()) bool {
	originalStdoutPrintf := stdoutPrintf
	defer func() {
		stdoutPrintf = originalStdoutPrintf
	}()
	stdoutPrintfCallCount := 0
	stdoutPrintf = func(string, ...interface{}) {
		stdoutPrintfCallCount++
	}
	whenThisIsCalled()
	return stdoutPrintfCallCount == 1
}
