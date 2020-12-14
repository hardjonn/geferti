package logging

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// Fields gives us a way to setup some extra JSON like fields
type Fields map[string]interface{}

// Logger represents common interface for logging function
type Logger interface {
	WithFields(fields *Fields) Logger
	ClearFields()

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

// SetLogger is the setter for the log variable, it should be the only way to assign value to log
func SetLogger(logger Logger) {
	Log = logger
}
