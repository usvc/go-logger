package logger

import (
	"path/filepath"
	"runtime"
)

func LogrusCallerPrettyfier(frame *runtime.Frame) (function string, file string) {
	return frame.Function, filepath.Base(frame.File)
}
