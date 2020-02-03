package logger

// Config holds configuration for a logger instance
type Config struct {
	Format         Format
	Level          Level
	Output         Output
	OutputFilePath string
	ReportCaller   bool
	Type           Type
}
