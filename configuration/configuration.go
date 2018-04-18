package configuration

import (
	"sync"

	"../storage"
)

type config struct {
	storage storage.Storager
}

var (
	cfg  *config
	once sync.Once
)

func get() *config {
	once.Do(func() {
		cfg = &config{}
		cfg.storage, _ = storage.Get("file", "./local/filestorage/")
	})
	return cfg
}

// GetStorage ...
func GetStorage() storage.Storager {
	return get().storage
}

// Close ...
func Close() {
	GetStorage().Close()
}
