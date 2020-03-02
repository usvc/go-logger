package logger

// Options holds configuration for a logger instance
type Options struct {
	Fields         map[string]interface{}
	Format         Format
	Level          Level
	Output         Output
	OutputFilePath string
	ReportCaller   bool
	Type           Type
}

func (opt *Options) AssignDefaults() {
	if !ValidFormats.Includes(opt.Format) {
		opt.Format = DefaultFormat
	}

	if !ValidLevels.Includes(opt.Level) {
		opt.Level = DefaultLevel
	}

	if !ValidTypes.Includes(opt.Type) {
		opt.Type = DefaultType
	}

	if !ValidOutputs.Includes(opt.Output) {
		opt.Output = DefaultOutput
	}
}
