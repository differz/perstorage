package repositories

import "../../core"

// OrderRepository ...
type OrderRepository interface {
	StoreOrder(order core.Order) (int, error)
	FindOrderByID(id int) (core.Order, bool)
	FindOrderByLink(link string) (core.Order, bool)
}
