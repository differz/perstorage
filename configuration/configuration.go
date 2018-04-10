package configuration

import (
	"database/sql"
	"sync"

	"../storage"
)

type Config struct {
	Storage    storage.Storager
	Connection *sql.DB
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		repo, _ := storage.Get("file", 1)

		cfg = &Config{}
		cfg.Storage = repo
		cfg.Connection = repo.InitDB()

		repo.Migrate(cfg.Connection)
	})
	return cfg
}

func (c Config) Close() {
	c.Connection.Close()
}
