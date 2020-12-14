package migrator

import (
	"github.com/hardjonn/geferti/migrations"
	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/migrator/commands"
	"github.com/hardjonn/geferti/pkg/storage/mysql"
	repository "github.com/hardjonn/geferti/pkg/storage/mysql/migration"
)

// Migrator defines the migrator struct.
type Migrator struct {
	config *config.C
	// logger zerolog.Logger
}

// New instantiates the new instance.
func New(config *config.C) (*Migrator, error) {
	// logger, err := logger.New(config.Logger)
	// if err != nil {
	// 	return nil, err
	// }

	return &Migrator{
		config: config,
		// logger: logger,
	}, nil
}

// Execute the cli command.
func (m *Migrator) Execute() error {
	// m.logger.Debug().Msg(m.config.DB.User)
	db, err := mysql.NewStorage(m.config.DB)
	if err != nil {
		return err
	}

	repo := repository.New(db)

	// migrator is server
	// commands are routes
	// each command should have access to the migrating service
	// migrating service has access to the repository
	// each migration should have access to the service as well

	migrator := migrations.NewService(repo, m.config.Migration)
	commands.Execute(migrator)

	return nil
}
