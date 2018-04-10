package repositories

import "../../core"

// ItemRepository ...
type ItemRepository interface {
	StoreItem(item core.Item) (int, error)
	FindItemByID(id int) (core.Item, bool)
}
