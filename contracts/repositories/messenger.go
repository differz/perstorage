package repositories

import "perstorage/core"

// Messenger repository
type Messenger interface {
	StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error
	FindCustomerChatID(customer core.Customer, messengerName string) (int, bool)
	IsRegisteredChatID(chatID int, messengerName string) (int, bool)
}
