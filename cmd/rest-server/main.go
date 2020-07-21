package main

import (
	"fmt"
	"geferti/pkg/config"
	"geferti/pkg/http/rest"
)

func main() {
	appConfig, err := config.New(".", ".env", "dotenv")
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}

	s, err := rest.New(appConfig)
	if err != nil {
		panic(fmt.Errorf("Could not initializing the REST server: %s", err))
	}

	if err := s.Start(); err != nil {
		panic(fmt.Errorf("Could not start the REST server: %s", err))
	}

	// logger, err := logger.New(appConfig.Logger)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error initializing logger: %s", err))
	// }

	// logger.Debug().Msg("app start")
	/*
		1. init db/store
		2. create api server
			- pass dependencies
				- logger
				- config
				- db/store
		3. start api server
			- before starting check if the server in the db
	*/
}
