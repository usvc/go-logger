package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type OptionsTests struct {
	testLogrusEntry *logrus.Entry
	suite.Suite
}

func TestOptions(t *testing.T) {
	suite.Run(t, &OptionsTests{})
}

func (s *OptionsTests) TestAssignDefaults() {
	options := Options{}
	options.AssignDefaults()
	s.Equal(DefaultFormat, options.Format)
	s.Equal(DefaultLevel, options.Level)
	s.Equal(DefaultOutput, options.Output)
	s.Equal(DefaultType, options.Type)
}
