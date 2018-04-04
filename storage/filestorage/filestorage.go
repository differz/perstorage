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

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) {

}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) core.Item {
	return core.Item{}
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(item core.Order) {

}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) core.Order {
	return core.Order{}
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) {

}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) core.Customer {
	return core.Customer{}
}

func init() {
	storage.Register("file", New())
}
