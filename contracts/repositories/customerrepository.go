package repositories

import "../../core"

// CustomerRepository ...
type CustomerRepository interface {
	StoreCustomer(item core.Customer) (int, error)
	FindCustomerByID(id int) (core.Customer, bool)
}
