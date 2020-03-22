package usecases

import (
	"perstorage/common"
	"perstorage/contracts/messengers"
	"perstorage/contracts/usecases"
	"perstorage/messenger"
	"perstorage/storage"
	uc "perstorage/usecases"
)

// CustomerMessengerUseCase object
type CustomerMessengerUseCase struct {
	repo storage.Storager
	msgr messenger.Messenger
}

// NewCustomerMessengerUseCase constructor
func NewCustomerMessengerUseCase(repo storage.Storager, msgr messenger.Messenger) CustomerMessengerUseCase {
	return CustomerMessengerUseCase{
		repo: repo,
		msgr: msgr,
	}
}

// ListenChat listen messengers chat
func (u CustomerMessengerUseCase) ListenChat() {
	request := messengers.ListenChatRequest{
		Repo: u.repo,
	}
	output := CustomerMessengerResponse{
		repo:         u.repo,
		msgr:         u.msgr,
		placeOrder:   uc.NewPlaceOrderUseCase(u.repo),
		orderMessage: uc.NewOrderMessageUseCase(u.repo, u.msgr),
	}
	go u.msgr.ListenChat(request, output)
}

// CustomerMessengerResponse object to response
type CustomerMessengerResponse struct {
	repo         storage.Storager
	msgr         messenger.Messenger
	placeOrder   contracts.PlaceOrderInput
	orderMessage contracts.OrderMessageInput
}

// OnResponse register new chatID to customer messenger
func (r CustomerMessengerResponse) OnResponse(request messengers.ListenChatRequest) {
	if common.ValidatePhone(request.Phone) {
		go r.registerMessenger(request.Phone, request.Messenger, request.ChatID)
	}
	if request.FileURL != "" {
		go r.downloadFile(request.FileURL, request.FileName, request.FileSize, request.Messenger, request.ChatID, request.Description)
	}
}
