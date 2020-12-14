package migration

import (
	"database/sql"
)

// Migration defines the storage form of a migration.
type Migration struct {
	Name       string       `db:"name"`
	MigratedAt sql.NullTime `db:"migrated_at"`
}
