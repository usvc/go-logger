package logger

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StdoutTests struct {
	suite.Suite
}

func TestStdout(t *testing.T) {
	suite.Run(t, &StdoutTests{})
}

func (s *StdoutTests) Test_print() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelTrace,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.print("hello")
	s.Equal("hello\n", output.String())
}

func (s *StdoutTests) Test_printf() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelTrace,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.printf("hello %s", "world")
	s.Equal("hello world\n", output.String())
}

func (s *StdoutTests) TestTracex() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelTrace,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.Trace("hey")
	stdout.Tracef("hello %s", "world")
	s.Equal("hey\nhello world\n", output.String())
	stdout.Level = LevelDebug
	stdout.Trace("bye")
	s.NotContains("bye", output.String())
}

func (s *StdoutTests) TestDebugx() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelDebug,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.Debug("hey")
	stdout.Debugf("hello %s", "world")
	s.Equal("hey\nhello world\n", output.String())
	stdout.Level = LevelInfo
	stdout.Debug("bye")
	s.NotContains("bye", output.String())
}

func (s *StdoutTests) TestInfox() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelInfo,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.Info("hey")
	stdout.Infof("hello %s", "world")
	s.Equal("hey\nhello world\n", output.String())
	stdout.Level = LevelWarn
	stdout.Info("bye")
	s.NotContains("bye", output.String())
}

func (s *StdoutTests) TestWarnx() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelWarn,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.Warn("hey")
	stdout.Warnf("hello %s", "world")
	s.Equal("hey\nhello world\n", output.String())
	stdout.Level = LevelError
	stdout.Warn("bye")
	s.NotContains("bye", output.String())
}

func (s *StdoutTests) TestErrorx() {
	var output bytes.Buffer
	stdout := Stdout{
		Level:        LevelError,
		OutputStream: bufio.NewWriter(&output),
	}
	stdout.Error("hey")
	stdout.Errorf("hello %s", "world")
	s.Equal("hey\nhello world\n", output.String())
}
