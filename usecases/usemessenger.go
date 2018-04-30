package usecases

import (
	"log"

	"../core"
	"../messenger"
	"../storage"
)

// CustomerMessengerUseCase object
type CustomerMessengerUseCase struct {
	repo storage.Storager
	msgr messenger.Messenger
	//
	subject     string
	description string
}

// NewCustomerMessengerUseCase constructor
func NewCustomerMessengerUseCase(repo storage.Storager, msgr messenger.Messenger) CustomerMessengerUseCase {
	return CustomerMessengerUseCase{
		repo:        repo,
		msgr:        msgr,
		description: "new customer messenger",
	}
}

// ListenChat listen messengers chat
func (u CustomerMessengerUseCase) ListenChat() {
	output := CustomerMessengerResponse{
		repo: u.repo,
	}
	go u.msgr.ListenChat(output)
}

// CustomerMessengerResponse object to response
type CustomerMessengerResponse struct {
	repo storage.Storager
}

// OnResponse register new chatID to customer messenger
func (r CustomerMessengerResponse) OnResponse(phone string, messengerName string, chatID int) {
	if validatePhone(phone) {
		go r.registerMessenger(phone, messengerName, chatID)
	}
}

func (r CustomerMessengerResponse) registerMessenger(phone string, messengerName string, chatID int) {
	customerID, err := core.GetCustomerIDByPhone(phone)
	if err != nil {
		log.Printf("can't get customer id by phone %s %e", phone, err)
		return
	}

	customer, ok := r.repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = phone
		r.repo.StoreCustomer(customer)
	}

	_, ok = r.repo.FindCustomerChatID(customer, messengerName)
	if !ok {
		r.repo.StoreCustomerMessenger(customer, messengerName, chatID)
	}
}

func validatePhone(phone string) bool {
	return phone != ""
}
