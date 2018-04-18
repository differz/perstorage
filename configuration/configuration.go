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
		cfg = &config{}
		cfg.storage, cfg.connection, _ = storage.Get("file", "./local/filestorage/")
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
