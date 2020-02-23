package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type LevelTests struct {
	testLogrusEntry *logrus.Entry
	suite.Suite
}

func TestLevel(t *testing.T) {
	suite.Run(t, &LevelTests{})
}

func (s *LevelTests) TestFormatsIncludes() {
	levels := Levels{LevelTrace, LevelWarn}
	s.True(levels.Includes(LevelTrace))
	s.True(levels.Includes(LevelWarn))
	s.False(levels.Includes(LevelInfo))
}
