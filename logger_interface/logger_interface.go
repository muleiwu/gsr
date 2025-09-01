package logger_interface

type LoggerFieldInterface interface {
	GetKey() string
	GetValue() any
}

type LoggerInterface interface {
	Debug(format string, args ...LoggerFieldInterface)
	Info(format string, args ...LoggerFieldInterface)
	Notice(format string, args ...LoggerFieldInterface)
	Error(format string, args ...LoggerFieldInterface)
	Warn(format string, args ...LoggerFieldInterface)
	Fatal(format string, args ...LoggerFieldInterface)
}
