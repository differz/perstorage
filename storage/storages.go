package storage

import (
	"fmt"
	"log"

	"../common"
	"../contracts/repositories"
)

// Storager consist of all repository contracts that each should implement
type Storager interface {
	repositories.Customer
	repositories.Order
	repositories.Item
	repositories.Messenger
	Init(args ...string)
	Close()
	//	String() string
}

const component = "storages"

var storages = make(map[string]Storager)

// Register messenger into factory by name
func Register(name string, storage Storager) {
	if storage == nil {
		log.Panicf("storage factory %s does not exist", name)
	}
	_, registered := storages[name]
	if registered {
		log.Printf("storage %s already registered", name)
	}
	storages[name] = storage
}

// Get messenger by name from factory
func Get(name string, args ...string) (Storager, error) {
	storage, ok := storages[name]
	if !ok {
		return nil, fmt.Errorf("unknown storage type: %s", name)
	}
	storage.Init(args...)
	return storage, nil
}

// String presents view of factory map
func String() string {
	return fmt.Sprint(storages)
}

// Print view of factory map to console
func Print() {
	common.ContextUpMessage(component, String())
}
