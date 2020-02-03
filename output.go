package logger

type Output string

const (
	OutputStdout     Output = "STDOUT"
	OutputStderr     Output = "STDERR"
	OutputFileSystem Output = "FS"
)
