package storage

import (
	"database/sql"
	"fmt"
	"log"

	"../contracts/repositories"
)

// Storager ...
type Storager interface {
	repositories.CustomerRepository
	repositories.OrderRepository
	repositories.ItemRepository
	InitDB() *sql.DB
	Migrate()
	//	String() string
}

var storages = make(map[string]Storager)

// Register ...
func Register(name string, storage Storager) {
	if storage == nil {
		log.Panicf("Storage factory %s does not exist", name)
	}
	_, registered := storages[name]
	if registered {
		log.Printf("Storage %s already registered", name)
	}
	storages[name] = storage
}

// Get ...
func Get(name string, x int) (Storager, error) {
	storage, ok := storages[name]
	if !ok {
		return nil, fmt.Errorf("Unknown storage type: %s", name)
	}
	return storage, nil
}

// Print ...
func Print() {
	fmt.Println(storages)
}
