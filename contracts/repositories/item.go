package repositories

import "../../core"

// Item repository
type Item interface {
	StoreItem(item core.Item) (int, error)
	FindItemByID(id int) (core.Item, bool)
}
