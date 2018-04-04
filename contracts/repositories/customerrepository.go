package repositories

import "../../core"

// CustomerRepository ...
type CustomerRepository interface {
	StoreCustomer(item core.Customer)
	FindCustomerById(id int) core.Customer
}
