package mysql

import "time"

// Node defines the storage form of a node.
type Node struct {
	ID        []byte    `db:"id"`
	UUID      string    `db:"uuid"`
	Key       string    `db:"key"`
	DeletedAt time.Time `db:"deleted_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}
