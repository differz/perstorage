package repositories

import "../../core"

// CustomerRepository ...
type CustomerRepository interface {
	StoreCustomer(customer core.Customer) (int, error)
	FindCustomerByID(id int) (core.Customer, bool)
}
