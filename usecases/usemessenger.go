package usecases

import (
	"strconv"
	"strings"

	"../messenger"
	"../storage"
)

// CustomerMessengerUseCase ...
type CustomerMessengerUseCase struct {
	repo storage.Storager
	msgr messenger.Messenger
	//
	subject     string
	description string
}

// NewCustomerMessengerUseCase ...
func NewCustomerMessengerUseCase(repo storage.Storager, msgr messenger.Messenger) CustomerMessengerUseCase {
	return CustomerMessengerUseCase{
		repo:        repo,
		msgr:        msgr,
		description: "new",
	}
}

// ListenChat ...
func (u CustomerMessengerUseCase) ListenChat() {
	output := CustomerMessengerResponse{
		repo: u.repo,
	}
	go u.msgr.ListenChat(output)
}

// CustomerMessengerResponse ...
type CustomerMessengerResponse struct {
	repo storage.Storager
}

// OnResponse ...
func (r CustomerMessengerResponse) OnResponse(phone string, messengerName string, chatID int) {
	if validatePhone(phone) {
		go r.registerMessenger(phone, messengerName, chatID)
	}
}

func (r CustomerMessengerResponse) registerMessenger(phone string, messengerName string, chatID int) {
	customerID, err := strconv.Atoi(strings.Replace(phone, "+", "", 1))
	if err != nil {
		// TODO error
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
