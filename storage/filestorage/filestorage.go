package filestorage

import (
	"../../core"
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
func (s Storage) StoreItem(item core.Item) {

}

// Get file from storage
func (s Storage) FindItemById(id int) {

}

// Save bucket to storage
func (s Storage) StoreOrder(item core.Item) {

}

// Get bucket from storage
func (s Storage) FindOrderById(id int) {

}

// Save client to storage
func (s Storage) StoreCustomer(item core.Item) {

}

// Get client from storage
func (s Storage) FindCustomerById(id int) {

}

func init() {
	storage.Register("file", New())
}
