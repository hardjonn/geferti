// Package factory handles creating concrete logger with factory method pattern
package factory

import (
	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/errs"
	"github.com/hardjonn/geferti/pkg/logging"
	"github.com/hardjonn/geferti/pkg/logging/zerolog"
)

// map logger code to logger builder
var logFactoryBuilderMap = map[string]LogFbInterface{
	config.ZEROLOG: &zerolog.Factory{},
}

// LogFbInterface is an interface for logger factory
type LogFbInterface interface {
	Build(*config.Logger) (logging.Logger, error)
}

// GetLogFactoryBuilder is an accessor for factoryBuilderMap
func GetLogFactoryBuilder(key string) (LogFbInterface, error) {
	if builder, ok := logFactoryBuilderMap[key]; ok {
		return builder, nil
	}

	return nil, errs.E(errs.Op("logging.factory.getBuilder"), errs.StatusNotFound, "logger factory builder not found")
}

// Build creates a logger instance
func Build(lc *config.Logger) (logging.Logger, error) {
	handler := lc.Handler

	logFactory, err := GetLogFactoryBuilder(handler)
	if err != nil {
		return nil, err
	}

	l, err := logFactory.Build(lc)
	if err != nil {
		return l, err
	}

	return l, nil
}
