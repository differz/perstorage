package config

import (
	"sync"

	"../storage"
)

type Config struct {
	Storage storage.Storager
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

	})
	return cfg
}
