package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type OutputTests struct {
	testLogrusEntry *logrus.Entry
	suite.Suite
}

func TestOutput(t *testing.T) {
	suite.Run(t, &OutputTests{})
}

func (s *OutputTests) TestOutputsIncludes() {
	outputs := Outputs{OutputStdout, OutputStderr}
	s.True(outputs.Includes(OutputStdout))
	s.True(outputs.Includes(OutputStderr))
	s.False(outputs.Includes(OutputFileSystem))
}
