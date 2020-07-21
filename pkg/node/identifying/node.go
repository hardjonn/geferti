package identifying

import (
	"database/sql"
)

// Node defines the storage form of a node.
type Node struct {
	ID        []byte       `db:"id"`
	UUID      string       `db:"uuid"`
	Key       string       `db:"key"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt sql.NullTime `db:"created_at"`
}
