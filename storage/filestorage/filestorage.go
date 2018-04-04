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
func (s Storage) FindItemById(id int) core.Item {
	return core.Item{}
}

// Save bucket to storage
func (s Storage) StoreOrder(item core.Order) {

}

// Get bucket from storage
func (s Storage) FindOrderById(id int) core.Order {
	return core.Order{}
}

// Save client to storage
func (s Storage) StoreCustomer(item core.Customer) {

}

// Get client from storage
func (s Storage) FindCustomerById(id int) core.Customer {
	return core.Customer{}
}

func init() {
	storage.Register("file", New())
}
