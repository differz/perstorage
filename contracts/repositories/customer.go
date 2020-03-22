package repositories

import "perstorage/core"

// Customer repository
type Customer interface {
	StoreCustomer(customer core.Customer) (int, error)
	FindCustomerByID(id int) (core.Customer, bool)
}
