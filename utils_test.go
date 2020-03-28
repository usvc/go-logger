package logger

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) TestNewLogrusEntry_withDefaults() {
	defer func() {
		if r := recover(); r != nil {
			s.True(false, r)
		}
	}()
	var _ Logger = NewLogrusEntry()
}

func (s *UtilsTests) TestNew_withDefaults() {
	defer func() {
		if r := recover(); r != nil {
			s.True(false, r)
		}
	}()
	var _ Logger = New()
}

func (s *UtilsTests) TestNew_TypeLevelled_withCustom() {
	var output bytes.Buffer
	var log = New(Options{
		Level:        LevelTrace,
		Type:         TypeLevelled,
		Output:       OutputCustom,
		OutputStream: &output,
	})
	log.Trace("hello world")
	s.Contains(output.String(), "@level=trace")
	s.Contains(output.String(), "@timestamp=20")
	s.Contains(output.String(), "@message=\"hello world\"")
}

func (s *UtilsTests) TestNew_TypeLevelled_withCustom_streamNotSpecified() {
	output := captureStdoutFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeLevelled,
			Output: OutputCustom,
		})
		log.Trace("hello world")
	})
	s.Contains(output, "@level=trace")
	s.Contains(output, "@timestamp=20")
	s.Contains(output, "@message=\"hello world\"")
}

func (s *UtilsTests) TestNew_TypeLevelled_withCustom_withFields() {
	var output bytes.Buffer
	var log = New(Options{
		Fields: map[string]interface{}{
			"module": "test utils",
		},
		Level:        LevelTrace,
		Type:         TypeLevelled,
		Output:       OutputCustom,
		OutputStream: &output,
	})
	log.Trace("hello world")
	s.Contains(output.String(), "@level=trace")
	s.Contains(output.String(), "@timestamp=20")
	s.Contains(output.String(), "@message=\"hello world\"")
	s.Contains(output.String(), "module=\"test utils\"")
}

func (s *UtilsTests) TestNew_TypeLevelled_withFileSystem() {
	logFilePath := "./.test_utils_new_type_levelled_with_file_system.log"
	file, err := os.Create(logFilePath)
	defer func() {
		err = file.Close()
		s.Nil(err)
	}()
	s.Nil(err)
	var log = New(Options{
		Level:          LevelTrace,
		Type:           TypeLevelled,
		Output:         OutputFileSystem,
		OutputFilePath: logFilePath,
	})
	log.Trace("hello world")
	output, err := ioutil.ReadAll(file)
	s.Nil(err)
	s.Contains(string(output), "@level=trace")
	s.Contains(string(output), "@timestamp=20")
	s.Contains(string(output), "@message=\"hello world\"")
	err = os.Remove(logFilePath)
	s.Nil(err)
}

func (s *UtilsTests) TestNew_TypeLevelled_withStder() {
	output := captureStderrFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeLevelled,
			Output: OutputStderr,
		})
		log.Trace("hello world")
	})
	s.Contains(output, "@level=trace")
	s.Contains(output, "@timestamp=20")
	s.Contains(output, "@message=\"hello world\"")
}

func (s *UtilsTests) TestNew_TypeLevelled_withStdout() {
	output := captureStdoutFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeLevelled,
			Output: OutputStdout,
		})
		log.Trace("hello world")
	})
	s.Contains(output, "@level=trace")
	s.Contains(output, "@timestamp=20")
	s.Contains(output, "@message=\"hello world\"")
}

func (s *UtilsTests) TestNew_TypeLevelled_withJSONFormat() {
	output := captureStdoutFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeLevelled,
			Format: FormatJSON,
		})
		log.Trace("hello world")
	})
	s.Contains(output, "\"@level\":\"trace\"")
	s.Contains(output, "\"@timestamp\":\"20")
	s.Contains(output, "\"@message\":\"hello world\"")
}

func (s *UtilsTests) TestNew_TypeNoOp() {
	var log Logger = New(Options{Type: TypeNoOp})
	s.NotNil(log)
}

func (s *UtilsTests) TestNew_TypeStdout_withCustom() {
	var output bytes.Buffer
	var log = New(Options{
		Level:        LevelTrace,
		Type:         TypeStdout,
		Output:       OutputCustom,
		OutputStream: &output,
	})
	log.Trace("hello world")
	s.Equal("hello world\n", output.String())
}

func (s *UtilsTests) TestNew_TypeStdout_withCustom_streamNotSpecified() {
	output := captureStdoutFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeStdout,
			Output: OutputCustom,
		})
		log.Trace("hello world")
	})
	s.Equal("hello world\n", output)
}

func (s *UtilsTests) TestNew_TypeStdout_withFileSystem() {
	logFilePath := "./.test_utils_new_type_stdout_with_file_system.log"
	file, err := os.Create(logFilePath)
	defer func() {
		err = file.Close()
		s.Nil(err)
	}()
	s.Nil(err)
	var log = New(Options{
		Level:          LevelTrace,
		Type:           TypeStdout,
		Output:         OutputFileSystem,
		OutputFilePath: logFilePath,
	})
	log.Trace("hello world")
	output, err := ioutil.ReadAll(file)
	s.Nil(err)
	s.Equal("hello world\n", string(output))
	err = os.Remove(logFilePath)
	s.Nil(err)
}

func (s *UtilsTests) TestNew_TypeStdout_withStdout() {
	output := captureStdoutFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeStdout,
			Output: OutputStdout,
		})
		log.Trace("hello world")
	})
	s.Equal("hello world\n", output)
}

func (s *UtilsTests) TestNew_TypeStdout_withStderr() {
	output := captureStderrFrom(func() {
		var log = New(Options{
			Level:  LevelTrace,
			Type:   TypeStdout,
			Output: OutputStderr,
		})
		log.Trace("hello world")
	})
	s.Equal("hello world\n", output)
}
