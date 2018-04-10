package repositories

import "../../core"

// OrderRepository ...
type OrderRepository interface {
	StoreOrder(item core.Order) (int, error)
	FindOrderByID(id int) (core.Order, bool)
}
