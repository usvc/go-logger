package logger

type Config struct {
	Format       Format
	Level        Level
	ReportCaller bool
	Type         Type
}
