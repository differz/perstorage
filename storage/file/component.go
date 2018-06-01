package file

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sync"

	"../../common"
	"../../configuration"
	"../../storage"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	// sqlite
	_ "github.com/mattes/migrate/source/file"
)

// Storage object for file db
type Storage struct {
	name       string
	dir        string
	connection *sql.DB
}

const component = "file"

var mutex sync.RWMutex

// New create storage instance
func New() *Storage {
	return &Storage{
		name: "file.db",
	}
}

// Init db and create connection. Do migration if needed.
func (s *Storage) Init(args ...string) {
	s.dir = args[0]
	common.ContextUpMessage(component, "init file storage on "+s.dir)
	err := os.MkdirAll(s.dir, os.ModePerm)
	if err != nil {
		log.Fatalf("can't create directory %s %e", s.dir, err)
	}

	mutex.Lock()
	defer mutex.Unlock()

	file := s.dir + "perstorage.db" + "?_busy_timeout=5000"
	s.connection, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("can't open sqlite storage %e", err)
	}
	s.connection.SetMaxOpenConns(1)
	path := "file://" + filepath.ToSlash(configuration.ExecutableDir()) + "/storage/file/migrations"
	s.migrate(path)
}

func (s Storage) migrate(loc string) {
	driver, err := sqlite3.WithInstance(s.connection, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("can't init sqlite driver %e", err)
	}
	m, err := migrate.NewWithDatabaseInstance(loc, "sqlite3", driver)
	if err != nil {
		log.Fatalf("can't get sqlite migration instance %e", err)
	}
	err = m.Up()
	if err == nil {
		common.ContextUpMessage("migrate", "database migrated successfully")
	} else if err == migrate.ErrNoChange {
		common.ContextUpMessage("migrate", "database is already up-to-date, no update required")
	} else {
		log.Fatalf("can't migrate database %e", err)
	}
}

// Close defer db.Close()
func (s Storage) Close() {
	s.connection.Close()
}

func (s Storage) String() string {
	return s.name
}

func init() {
	storage.Register(component, New())
}
