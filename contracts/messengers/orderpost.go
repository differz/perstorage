package messengers

// OrderPostInput ...
type OrderPostInput interface {
	ShowOrder(chatID int, message string) error
}
