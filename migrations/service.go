package migrations

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"os"
	"sort"
	"time"

	"github.com/hardjonn/geferti/pkg/config"
	repository "github.com/hardjonn/geferti/pkg/storage/mysql/migration"

	"github.com/iancoleman/strcase"
)

// ErrNotFound is used when a node could not be found
var ErrNotFound = errors.New("node not found")

// Service provides migration operations
type Service interface {
	Make(string) error
	Up(step int) error
	Down() error
	Init() error
}

// RepositoryInterface provides access to the migration repository
type RepositoryInterface interface {
	// InitMigrationSchema generate the initial migration schema
	InitMigrationSchema() error

	// GetAllMigrations fetches all the migrations from the database.
	GetAllMigrations() ([]repository.Migration, error)

	// SaveMigration persists the migration into the database.
	SaveMigration(string) error

	// GetStorage provides access to the db interface.
	GetStorage() repository.DBInterface
}

type migrationName struct {
	db     string
	file   string
	method string
}

// Migration describes a migration of the migration list.
type Migration struct {
	Name       string
	Up         func(RepositoryInterface) error
	Down       func(RepositoryInterface) error
	MigratedAt sql.NullTime
}

// Migrations is a simple store of all the available migrations.
type Migrations struct {
	ids   []string
	items map[string]*Migration
}

type service struct {
	r RepositoryInterface
	c *config.Migration
}

var migrations = &Migrations{
	ids:   []string{},
	items: map[string]*Migration{},
}

// NewService creates an migrating service with the necessary dependencies.
func NewService(r RepositoryInterface, c *config.Migration) Service {
	return &service{r, c}
}

// Make generates a new migration from the template with the given name.
func (s *service) Make(name string) error {
	fmt.Println("migrator MAKE")
	fmt.Println(name)

	n := makeName(name)
	fmt.Println(n)

	fmt.Println(s.c.Path)
	fmt.Println(s.c.Stub)

	in := struct {
		Name   string
		DbName string
	}{Name: n.method, DbName: n.db}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles(s.c.Stub))

	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template:" + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.go", s.c.Path, n.file))
	if err != nil {
		return errors.New("Unable to create migration file:" + err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file:" + err.Error())
	}

	fmt.Println("Generated new migration files...", f.Name())

	return nil
}

// Up migrates a migration and persists it to the database.
func (s *service) Up(step int) error {
	s.Init()

	count := 0
	// step = step | len(migration.ids)
	// for id := 0; id < len(migrations.ids) && step < count; id++
	for _, id := range migrations.ids {
		// stop if migration the specified amount of items
		if step > 0 && count == step {
			break
		}

		mg := migrations.items[id]

		// skip already migrated items
		if mg.MigratedAt.Valid {
			continue
		}

		// run the internal migration Up
		if err := mg.Up(s.r); err != nil {
			return err
		}

		// and mark the migration as successfully migrated by persisting into the db
		if err := s.r.SaveMigration(id); err != nil {
			return err
		}

		count++
	}

	fmt.Println("IDS")
	fmt.Println(migrations.ids)

	return nil
}

// Down rollbacks the migration.
func (s *service) Down() error {
	// for _, v := range migrations {
	// 	v.Down(s)
	// }

	return nil
}

// Init the migrations table and sync the migrated records
func (s *service) Init() error {
	err := s.r.InitMigrationSchema()
	if err != nil {
		return err
	}

	mm, err := s.r.GetAllMigrations()
	if err != nil {
		return err
	}

	migrations.items = syncMigratedItems(migrations.items, mm)

	return nil
}

// sync the state of the migrated items
func syncMigratedItems(items map[string]*Migration, mm []repository.Migration) map[string]*Migration {
	for _, m := range mm {
		if items[m.Name] != nil {
			items[m.Name].MigratedAt = m.MigratedAt
		}
	}

	return items
}

func makeName(name string) *migrationName {
	t := timestamp(time.Now())
	m := strcase.ToCamel(name)
	f := strcase.ToSnake(name)
	d := f

	return &migrationName{
		db:     fmt.Sprintf("%s_%s", t, d),
		file:   fmt.Sprintf("%s_%s", t, f),
		method: fmt.Sprintf("%s_%s", t, m),
	}
}

func timestamp(t time.Time) string {
	return fmt.Sprintf(
		"%d_%02d_%02d_%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Unix(),
	)
}

// push migration into the internal storage
func push(mg *Migration) {
	// put the actual migration info into the hash map storage
	migrations.items[mg.Name] = mg
	fmt.Println(migrations)

	// add the migration name/key into the list of ids in the ASC order
	migrations.ids = sortedInsert(migrations.ids, mg.Name)
}

// insert the given name/key into the sorted list of ids in the ASC order
func sortedInsert(ids []string, name string) []string {
	i := sort.Search(len(ids), func(i int) bool {
		return ids[i] > name
	})

	ids = append(ids, name)  // make some extra space at the end
	copy(ids[i+1:], ids[i:]) // shift elements
	ids[i] = name            // insert the element into the proper position

	return ids
}
