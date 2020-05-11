package logger

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type FormatTests struct {
	testLogrusEntry *logrus.Entry
	suite.Suite
}

func TestFormat(t *testing.T) {
	suite.Run(t, &FormatTests{})
}

func (s *FormatTests) SetupTest() {
	timestamp, err := time.Parse(TimestampFormat, "20200208134001")
	if err != nil {
		panic(err)
	}
	s.testLogrusEntry = &logrus.Entry{
		Data: logrus.Fields{
			"hello":   "world",
			"number":  1,
			"float":   3.142,
			"boolean": true,
		},
		Level:   logrus.DebugLevel,
		Message: "message",
		Time:    timestamp,
	}
}

func (s *FormatTests) TestFormatsIncludes() {
	formats := Formats{FormatJSON}
	s.True(formats.Includes(FormatJSON))
	s.False(formats.Includes(FormatText))
}

func (s *FormatTests) TestLogrusCallerPrettyfier() {
	frame := &runtime.Frame{
		Function: "function",
		File:     "/path/to/some/file",
	}
	function, file := LogrusCallerPrettyfier(frame)
	s.Equal("function", function)
	s.Equal("file", file)
}

func (s *FormatTests) TestLogrusCallerPrettyfier_withFunctionPath() {
	frame := &runtime.Frame{
		Function: "path/to/function",
		File:     "/path/to/some/file",
	}
	function, file := LogrusCallerPrettyfier(frame)
	s.Equal("function", function)
	s.Equal("file", file)
}

func (s *FormatTests) TestFormatJSONPreset() {
	var formatter logrus.Formatter = FormatJSONPreset
	formattedText, err := formatter.Format(s.testLogrusEntry)
	if err != nil {
		panic(err)
	}
	s.Equal(
		fmt.Sprintf(
			`{"%s":{"boolean":true,"float":3.142,"hello":"world","number":1}`+
				`,"%s":"debug"`+
				`,"%s":"message"`+
				`,"%s":"20200208134001"}%s`,
			FieldData,
			FieldLevel,
			FieldMessage,
			FieldTimestamp,
			"\n",
		),
		string(formattedText),
	)
}

func (s *FormatTests) TestFormatTextPreset() {
	var formatter logrus.Formatter = FormatTextPreset
	formattedText, err := formatter.Format(s.testLogrusEntry)
	if err != nil {
		panic(err)
	}
	formattedTextAsString := string(formattedText)
	s.Contains(formattedTextAsString,
		fmt.Sprintf("%s=debug", FieldLevel))
	s.Contains(formattedTextAsString,
		fmt.Sprintf("%s=message", FieldMessage))
	s.Contains(formattedTextAsString,
		"boolean=true")
	s.Contains(formattedTextAsString,
		"hello=world")
	s.Contains(formattedTextAsString,
		"number=1")
	s.Contains(formattedTextAsString,
		"float=3.142")
}
