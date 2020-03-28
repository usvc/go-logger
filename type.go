package logger

// Type represents a type of logger
type Type string

// Types is a slice of Types
type Types []Type

// Includes returns true if the provided :typ is a proper log type
func (types Types) Includes(this Type) bool {
	for _, t := range types {
		if t == this {
			return true
		}
	}
	return false
}

const (
	TypeLevelled Type = "levelled"
	TypeNoOp     Type = "noop"
	TypeStdout   Type = "stdout"

	DefaultType = TypeLevelled
)

var ValidTypes = Types{TypeLevelled, TypeNoOp, TypeStdout}
