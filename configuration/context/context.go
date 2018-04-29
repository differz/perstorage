package context

import (
	"fmt"
	"sync"

	"../../common"
	"../../configuration"
	"../../messenger"
	"../../storage"
)

type context struct {
	name      string
	storage   storage.Storager
	messenger messenger.Messenger
}

const component = "configuration"

var (
	ctx  *context
	once sync.Once
)

func getContext() *context {
	once.Do(func() {
		cfg := configuration.Get()
		common.ContextUpMessage(component, fmt.Sprint(cfg))

		ctx = &context{}
		ctx.name = "main"
		ctx.storage, _ = storage.Get(cfg.StorageName, cfg.StorageArgs)
		ctx.messenger, _ = messenger.Get(cfg.MessengerName, cfg.MessengerKey)
	})
	return ctx
}

// Name of current configuration
func Name() string {
	return getContext().name
}

// Messenger takes messenger @Bean
func Messenger() messenger.Messenger {
	return getContext().messenger
}

// Storage take storage @Bean
func Storage() storage.Storager {
	return getContext().storage
}

// Close all connections before exit
func Close() {
	Storage().Close()
}
