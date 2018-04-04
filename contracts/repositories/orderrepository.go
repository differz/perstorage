package repositories

import "../../core"

// OrderRepository ...
type OrderRepository interface {
	StoreOrder(item core.Order)
	FindOrderById(id int) core.Order
}
