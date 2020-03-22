package repositories

import "perstorage/core"

// Order repository
type Order interface {
	StoreOrder(order core.Order) (int, error)
	FindOrderByID(id int) (core.Order, bool)
	FindOrderByLink(link string) (core.Order, bool)
	GetOrders(strategy func() string) ([]core.Order, error)
	GetOrderedItems(order core.Order) ([]core.Item, bool)
	DeleteOrder(order core.Order) bool
}
