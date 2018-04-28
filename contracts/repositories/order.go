package repositories

import "../../core"

// Order repository
type Order interface {
	StoreOrder(order core.Order) (int, error)
	FindOrderByID(id int) (core.Order, bool)
	FindOrderByLink(link string) (core.Order, bool)
}
