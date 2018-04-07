package mongostorage

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
		name: "MongoDB",
	}
}

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) {

}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(item core.Order) {

}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
	return core.Order{}, false
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) {

}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
	return core.Customer{}, false
}

func init() {
	storage.Register("mongo", New())
}
