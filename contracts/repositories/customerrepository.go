package repositories

import "../../core"

// CustomerRepository ...
type CustomerRepository interface {
	StoreCustomer(item core.Customer)
	FindCustomerByID(id int) core.Customer
}
