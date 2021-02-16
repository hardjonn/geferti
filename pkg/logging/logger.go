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

// Debug is a short notation helper function to avoid stators: logging.Log.Debug -> logging.Debug
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Debugf is a short notation helper function to avoid stators: logging.Log.Debugf -> logging.Debugf
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

// Info is a short notation helper function to avoid stators: logging.Log.Info -> logging.Info
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Infof is a short notation helper function to avoid stators: logging.Log.Infof -> logging.Infof
func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

// Warn is a short notation helper function to avoid stators: logging.Log.Warn -> logging.Warn
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Warnf is a short notation helper function to avoid stators: logging.Log.Warnf -> logging.Warnf
func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

// Error is a short notation helper function to avoid stators: logging.Log.Error -> logging.Error
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Errorf is a short notation helper function to avoid stators: logging.Log.Errorf -> logging.Errorf
func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

// Fatal is a short notation helper function to avoid stators: logging.Log.Fatal -> logging.Fatal
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

// Fatalf is a short notation helper function to avoid stators: logging.Log.Fatalf -> logging.Fatalf
func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

// Trace is a short notation helper function to avoid stators: logging.Log.Trace -> logging.Trace
func Trace(args ...interface{}) {
	Log.Trace(args...)
}

// Tracef is a short notation helper function to avoid stators: logging.Log.Tracef -> logging.Tracef
func Tracef(format string, args ...interface{}) {
	Log.Tracef(format, args...)
}

// Panic is a short notation helper function to avoid stators: logging.Log.Panic -> logging.Panic
func Panic(args ...interface{}) {
	Log.Panic(args...)
}

// Panicf is a short notation helper function to avoid stators: logging.Log.Panicf -> logging.Panicf
func Panicf(format string, args ...interface{}) {
	Log.Panicf(format, args...)
}
