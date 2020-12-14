package mysql

import (
	"fmt"

	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/errs"

	// import it here since we do not need in the main package
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// NewStorage return a new MySQL storage
func NewStorage(config *config.DB) (*sqlx.DB, error) {
	conn := connectionString(config)

	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		return nil, errs.E(errs.Op("storage.mysql.migrations.NewStorage"), err, errs.StatusInternal)
	}

	return db, nil
}

func connectionString(config *config.DB) string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
