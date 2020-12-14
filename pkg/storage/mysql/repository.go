package mysql

import (
	"database/sql"

	"github.com/hardjonn/geferti/pkg/errs"
	"github.com/hardjonn/geferti/pkg/node/identifying"

	"github.com/jmoiron/sqlx"
)

// Storage stores data in MySQL database
type Storage struct {
	db *sqlx.DB
}

// New creates a new repository
func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// GetNodeByKey searches a node by the given key.
func (s *Storage) GetNodeByKey(key string) (identifying.Node, error) {
	n := identifying.Node{}

	err := s.db.Get(&n, "SELECT * FROM nodes n WHERE n.deleted_at IS NULL AND n.key = ? LIMIT 1", key)
	if err == sql.ErrNoRows {
		return n, errs.E(errs.Op("storage.mysql.GetNodeByKey"), identifying.ErrNotFound, errs.StatusNotFound)
	}

	return n, errs.E(errs.Op("storage.mysql.GetNodeByKey"), err, errs.StatusIO)
}

// AddNode adds a node to the storage.
func (s *Storage) AddNode(n identifying.Node) error {
	return nil
}

// ExecuteStatement executes an SQL statement
func (s *Storage) ExecuteStatement() error {
	return nil
}

// InitMigrationSchema persists the migration schema into the database.
func (s *Storage) InitMigrationSchema() error {
	return nil
}
