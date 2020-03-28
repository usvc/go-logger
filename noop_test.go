package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NoOpTests struct {
	suite.Suite
}

func TestNoOp(t *testing.T) {
	suite.Run(t, &NoOpTests{})
}

func (s *NoOpTests) Test_interfaceMatches() {
	var log Logger = &NoOp{}
	log.Trace("test_noop_Trace")
	log.Tracef("test_noop_Tracef")
	log.Debug("test_noop_Debug")
	log.Debugf("test_noop_Debugf")
	log.Info("test_noop_Info")
	log.Infof("test_noop_Infof")
	log.Warn("test_noop_Warn")
	log.Warnf("test_noop_Warnf")
	log.Error("test_noop_Error")
	log.Errorf("test_noop_Errorf")
}
