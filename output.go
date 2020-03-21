package logger

import (
	"fmt"
	"os"
)

// Output is a type of log output
type Output string

// Outputs is a slice of Outputs
type Outputs []Output

// Includes returns true if the provided Output :this is a proper log output
func (outputs Outputs) Includes(this Output) bool {
	for _, o := range outputs {
		if o == this {
			return true
		}
	}
	return false
}

const (
	OutputCustom     Output = "custom"
	OutputFileSystem Output = "fs"
	OutputStdout     Output = "stdout"
	OutputStderr     Output = "stderr"

	OutputFileSystemFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	OutputFileSystemMode  = os.ModePerm

	DefaultOutput = OutputStdout
)

var (
	DefaultOutputFilePath = fmt.Sprintf("./%s.log", RuntimeTimestamp)
	DefaultOutputStream   = &os.Stdout
)

var ValidOutputs = Outputs{
	OutputCustom,
	OutputFileSystem,
	OutputStdout,
	OutputStderr,
}
