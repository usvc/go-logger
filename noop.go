package logger

type NoOp struct{}

func (noop NoOp) Trace(_ ...interface{})            {}
func (noop NoOp) Tracef(_ string, _ ...interface{}) {}
func (noop NoOp) Debug(_ ...interface{})            {}
func (noop NoOp) Debugf(_ string, _ ...interface{}) {}
func (noop NoOp) Info(_ ...interface{})             {}
func (noop NoOp) Infof(_ string, _ ...interface{})  {}
func (noop NoOp) Warn(_ ...interface{})             {}
func (noop NoOp) Warnf(_ string, _ ...interface{})  {}
func (noop NoOp) Error(_ ...interface{})            {}
func (noop NoOp) Errorf(_ string, _ ...interface{}) {}
