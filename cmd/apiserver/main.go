package main

import (
	"fmt"
	"geferti/pkg/config"
	"geferti/pkg/logger"
)

func main() {
	appConfig, err := config.New(".", ".env", "dotenv")
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}

	logger, err := logger.New(appConfig.Logger)
	if err != nil {
		panic(fmt.Errorf("Fatal error initializing logger: %s", err))
	}

	logger.Debug().Msg("app start")
}
