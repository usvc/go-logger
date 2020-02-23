package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type TypeTests struct {
	testLogrusEntry *logrus.Entry
	suite.Suite
}

func TestType(t *testing.T) {
	suite.Run(t, &TypeTests{})
}

func (s *TypeTests) TestTypesIncludes() {
	types := Types{TypeStdout}
	s.True(types.Includes(TypeStdout))
	s.False(types.Includes(TypeLevelled))
}
