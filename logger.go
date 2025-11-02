package gsr

type Logger interface {
	Debug(format string, args ...LoggerField)
	Info(format string, args ...LoggerField)
	Notice(format string, args ...LoggerField)
	Error(format string, args ...LoggerField)
	Warn(format string, args ...LoggerField)
	Fatal(format string, args ...LoggerField)
}

type LoggerField interface {
	GetKey() string
	GetValue() any
}
