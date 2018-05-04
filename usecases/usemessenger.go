package usecases

import (
	"log"

	"../contracts/messengers"
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
	request := messengers.ListenChatRequest{
		Repo: u.repo,
	}
	output := CustomerMessengerResponse{
		repo: u.repo,
		msgr: u.msgr,
	}
	go u.msgr.ListenChat(request, output)
}

// CustomerMessengerResponse object to response
type CustomerMessengerResponse struct {
	repo storage.Storager
	msgr messenger.Messenger
}

// OnResponse register new chatID to customer messenger
func (r CustomerMessengerResponse) OnResponse(request messengers.ListenChatRequest) {
	if validatePhone(request.Phone) {
		go r.registerMessenger(request.Phone, request.Messenger, request.ChatID)
	}
	if request.FileID != "" {
		go r.downloadFile(request.FileID, request.FileName, request.FileSize)
	}
}

func (r CustomerMessengerResponse) registerMessenger(phone, messenger string, chatID int) {
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

	_, ok = r.repo.FindCustomerChatID(customer, messenger)
	if !ok {
		r.repo.StoreCustomerMessenger(customer, messenger, chatID)
	}
}

func (r CustomerMessengerResponse) downloadFile(id, name string, size int) {
	r.msgr.DownloadFile(id)
}

func validatePhone(phone string) bool {
	return phone != ""
}
