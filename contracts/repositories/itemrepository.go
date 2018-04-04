package repositories

import "../../core"

// ItemRepository ...
type ItemRepository interface {
	StoreItem(item core.Item)
	FindItemByID(id int) core.Item
}
