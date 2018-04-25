package repositories

import "../../core"

// MessengerRepository ...
type MessengerRepository interface {
	StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error
	FindCustomerChatID(customer core.Customer, messengerName string) (int, bool)
}
