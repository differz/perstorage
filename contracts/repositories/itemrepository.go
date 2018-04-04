package repositories

import "../../core"

// ItemRepository ...
type ItemRepository interface {
	StoreItem(item core.Item)
	FindItemById(id int) core.Item
}
