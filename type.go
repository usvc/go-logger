package logger

type Type string

const (
	TypeStdout   Type = "STDOUT"
	TypeLevelled Type = "LEVELLED"
	TypeText     Type = "TEXT"
)
