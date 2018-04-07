package repositories

import "../../core"

// OrderRepository ...
type OrderRepository interface {
	StoreOrder(item core.Order)
	FindOrderByID(id int) (core.Order, bool)
}
