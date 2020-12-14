package main

import (
	"fmt"

	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/logging"
	"github.com/hardjonn/geferti/pkg/logging/factory"
	"github.com/hardjonn/geferti/pkg/migrator"
)

func main() {
	appConfig, err := config.New(".", ".env", "dotenv")
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}

	// init log
	log, err := factory.Build(appConfig.Logger)
	logging.SetLogger(log)
	logging.Log.Debug("hello")
	logging.Log.WithFields(&logging.Fields{"a": "abc"}).Debug("hello world")
	logging.Log.WithFields(&logging.Fields{"int": 10}).Info("with int fields")
	logging.Log.Debug("without fields")

	m, err := migrator.New(appConfig)
	if err != nil {
		panic(fmt.Errorf("Could not initializing the Migrator: %s", err))
	}

	if err := m.Execute(); err != nil {
		panic(fmt.Errorf("Could not execute the migration command: %s", err))
	}
}
