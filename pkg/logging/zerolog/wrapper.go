package zerolog

import (
	"github.com/hardjonn/geferti/pkg/logging"
	"github.com/rs/zerolog"
)

type loggerWrapper struct {
	lw zerolog.Logger
	f  *logging.Fields
}

// WithFields setup JSON fields for the current output and returns the logger
func (logger *loggerWrapper) WithFields(fields *logging.Fields) logging.Logger {
	logger.f = fields
	return logger
}

// ClearFields clears the current fields set
func (logger *loggerWrapper) ClearFields() {
	logger.f = &logging.Fields{}
}

// Debug logs a debug level message
func (logger *loggerWrapper) Debug(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Debug().Fields(*logger.f).Msg(args[0].(string))
}

// Debugf logs a formatted debug level message
func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Debug().Fields(*logger.f).Msgf(format, args)
}

// Info logs an info level message
func (logger *loggerWrapper) Info(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Info().Fields(*logger.f).Msg(args[0].(string))
}

// Infof logs a formatted info level message
func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Info().Fields(*logger.f).Msgf(format, args)
}

// Warn logs a warn level message
func (logger *loggerWrapper) Warn(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Warn().Fields(*logger.f).Msg(args[0].(string))
}

// Warnf logs a formatted warn level message
func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Warn().Fields(*logger.f).Msgf(format, args)
}

// Error logs an error level message
func (logger *loggerWrapper) Error(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Error().Fields(*logger.f).Msg(args[0].(string))
}

// Errorf logs a formatted error level message
func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Error().Fields(*logger.f).Msgf(format, args)
}

// Fatal logs a fatal level message
func (logger *loggerWrapper) Fatal(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Fatal().Fields(*logger.f).Msg(args[0].(string))
}

// Fatalf logs a formatted fatal level message
func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Fatal().Fields(*logger.f).Msgf(format, args)
}

// Trace logs a trace level message
func (logger *loggerWrapper) Trace(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Trace().Fields(*logger.f).Msg(args[0].(string))
}

// Tracef logs a formatted trace level message
func (logger *loggerWrapper) Tracef(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Trace().Fields(*logger.f).Msgf(format, args)
}

// Panic logs a panic level message
func (logger *loggerWrapper) Panic(args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Panic().Fields(*logger.f).Msg(args[0].(string))
}

// Panicf logs a formatted panic level message
func (logger *loggerWrapper) Panicf(format string, args ...interface{}) {
	defer logger.ClearFields()
	logger.lw.Panic().Fields(*logger.f).Msgf(format, args)
}
