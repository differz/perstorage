package configuration

import (
	"database/sql"
	"sync"

	"../storage"
)

type config struct {
	storage    storage.Storager
	connection *sql.DB
}

var (
	cfg  *config
	once sync.Once
)

func get() *config {
	once.Do(func() {
		repo, _ := storage.Get("file", 1)

		cfg = &config{}
		cfg.storage = repo
		cfg.connection = repo.InitDB()
		repo.Migrate()
	})
	return cfg
}

// GetStorage ...
func GetStorage() storage.Storager {
	return get().storage
}

// Close ...
func Close() {
	get().connection.Close()
}
