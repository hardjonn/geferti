package rest

import (
	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/errs"
	"github.com/hardjonn/geferti/pkg/hosting"
	"github.com/hardjonn/geferti/pkg/logger"
	"github.com/hardjonn/geferti/pkg/node/identifying"
	"github.com/hardjonn/geferti/pkg/platform"
	"github.com/hardjonn/geferti/pkg/storage/mysql"

	"github.com/rs/zerolog"
)

// Server defines the server struct.
type Server struct {
	config *config.C
	logger zerolog.Logger
}

// New instantiates the new instance.
func New(config *config.C) (*Server, error) {
	logger, err := logger.New(config.Logger)
	if err != nil {
		return nil, err
	}

	return &Server{
		config: config,
		logger: logger,
	}, nil
}

// Start the server.
func (s *Server) Start() error {
	storage, err := mysql.NewStorage(s.config.DB)
	if err != nil {
		return err
	}

	host := platform.NewHost()
	hoster := hosting.NewService(host)

	mID, err := hoster.GetMachineID(s.config.App.Key)
	if err != nil {
		return errs.E(errs.Op("rest.server.start.GetMachineID"), err)
	}

	nodeIdentifier := identifying.NewService(storage)
	err = nodeIdentifier.IdentifyNode(mID)

	/*
		1. init storage
		2. init services
		3. register node
		4. init router
		5. start server
	*/

	return nil
}
