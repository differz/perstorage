package messengers

import "../../core"

// OrderPost ...
type OrderPost interface {
	ShowOrder(item core.Order) error
}
