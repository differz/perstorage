package repositories

import "../../core"

// Messenger repository
type Messenger interface {
	StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error
	FindCustomerChatID(customer core.Customer, messengerName string) (int, bool)
}
