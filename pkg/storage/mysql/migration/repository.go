package migration

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// DBInterface provides access to the underlying database instance.
type DBInterface interface {
	DB() *sqlx.DB
}

// Storage stores data in MySQL database
type Storage struct {
	db *sqlx.DB
}

// New creates a new repository
func New(db *sqlx.DB) *Storage {
	fmt.Println(db)
	return &Storage{db: db}
}

func newNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// InitMigrationSchema persists the migration schema into the database.
func (s *Storage) InitMigrationSchema() error {
	if _, err := s.db.Exec(schemaMigrations); err != nil {
		return err
	}

	return nil
}

// GetAllMigrations retrieves all the records from the migration table.
func (s *Storage) GetAllMigrations() ([]Migration, error) {
	migrations := []Migration{}
	if err := s.db.Select(&migrations, "SELECT * FROM migrations ORDER BY name ASC"); err != nil {
		return nil, err
	}

	return migrations, nil
}

// SaveMigration persists the migration into the database.
func (s *Storage) SaveMigration(name string) error {
	tx, err := s.db.Beginx()

	if err != nil {
		return err
	}

	t := newNullTime(time.Now())

	tx.NamedExec(
		"INSERT INTO migrations (name, migrated_at) VALUES (:name, :migrated_at)",
		&Migration{name, t},
	)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetStorage returns the storage instance as an interface to get an access to the underlying database.
func (s *Storage) GetStorage() DBInterface {
	return s
}

// DB returns the database instance.
func (s *Storage) DB() *sqlx.DB {
	return s.db
}
