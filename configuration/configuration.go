package configuration

import (
	"sync"

	"../messenger"
	"../storage"
)

type config struct {
	storage   storage.Storager
	messenger messenger.Messenger
}

var (
	cfg  *config
	once sync.Once
)

func get() *config {
	once.Do(func() {
		cfg = &config{}
		cfg.storage, _ = storage.Get("file", "./local/filestorage/")
		cfg.messenger, _ = messenger.Get("telegram", "529441026:AAEVlmwD87qxmP-dLsu5EwFovHVyKi2iVfE22")
	})
	return cfg
}

// GetMessenger ...
func GetMessenger() messenger.Messenger {
	return get().messenger
}

// GetStorage ...
func GetStorage() storage.Storager {
	return get().storage
}

// Close ...
func Close() {
	GetStorage().Close()
}
