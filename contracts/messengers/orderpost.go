package messengers

// OrderPostInput order to chatID contract
type OrderPostInput interface {
	ShowOrder(chatID int, message, description string) error
}
