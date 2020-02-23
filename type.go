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
	TypeStdout   Type = "stdout"
	TypeLevelled Type = "levelled"

	DefaultType = TypeLevelled
)

var ValidTypes = Types{TypeStdout, TypeLevelled}
