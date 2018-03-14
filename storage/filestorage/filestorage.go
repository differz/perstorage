package filestorage

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
		name: "file.txt",
	}
}

// Save file to storage
func (s Storage) Save() {

}

// Get file from storage
func (s Storage) Get() {

}

func init() {
	storage.Register("file", New())
}
