package mysql

import (
	"database/sql"
	"fmt"
	"geferti/pkg/config"
	"geferti/pkg/node/identifying"

	// import it here since we do not need in the main package
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// Storage stores data in MySQL database
type Storage struct {
	db *sqlx.DB
}

// NewStorage return a new MySQL storage
func NewStorage(config *config.DB) (*Storage, error) {
	conn := connectionString(config)

	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
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

// GetNodeByKey searches a node by the given key.
func (s *Storage) GetNodeByKey(key string) (identifying.Node, error) {
	n := identifying.Node{}

	err := s.db.Get(&n, "SELECT * FROM nodes n WHERE n.deleted_at IS NULL AND n.key = ? LIMIT 1", key)
	if err == sql.ErrNoRows {
		return n, identifying.ErrNotFound
	}

	return n, err
}
