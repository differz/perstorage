package mongostorage

import (
	"../../storage"
)

// Storage ...
type Storage struct {
	name string
}

// New create storage instance
func New() Storage {
	return Storage{
		name: "MongoDB",
	}
}

// Save file to storage
func (s Storage) Save() {

}

// Get file from storage
func (s Storage) Get() {

}

func init() {
	storage.Register("mongo", New())
}
