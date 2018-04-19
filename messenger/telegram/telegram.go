package telegram

import (
	"../../core"
	"../../messenger"
)

// Messenge ...
type Messenge struct {
	name string
}

// New create storage instance
func New() Messenge {
	return Messenge{
		name: "telegram",
	}
}

// ShowOrder ...
func (m Messenge) ShowOrder(item core.Order) error {

	return nil
}

func init() {
	messenger.Register("telegram", New())
}
