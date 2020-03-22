package context

import (
	"fmt"
	"sync"

	"perstorage/common"
	"perstorage/configuration"
	"perstorage/messenger"
	"perstorage/storage"
)

type context struct {
	name      string
	storage   storage.Storager
	messenger messenger.Messenger
}

const component = "context"

var (
	ctx  *context
	once sync.Once
)

func get() *context {
	once.Do(func() {
		cfg := configuration.Get()
		common.ContextUpMessage(configuration.Component(), fmt.Sprint(cfg))

		ctx = &context{}
		ctx.name = configuration.Name()
		ctx.storage, _ = storage.Get(cfg.StorageName, cfg.StorageArgs)
		ctx.messenger, _ = messenger.Get(cfg.MessengerName, cfg.MessengerKey, configuration.ServerAddress())
		common.ContextUpMessage(component, fmt.Sprint(ctx))
	})
	return ctx
}

// Name of current configuration
func Name() string {
	return get().name
}

// Messenger takes messenger @Bean
func Messenger() messenger.Messenger {
	return get().messenger
}

// Storage take storage @Bean
func Storage() storage.Storager {
	return get().storage
}

// Close all connections before exit
func Close() {
	Storage().Close()
}
