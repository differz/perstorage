package messengers

import (
	"../../configuration/context"
	"../../contracts/usecases"
	"../../usecases"
)

// Service object
type Service struct {
	customerMessenger contracts.CustomerMessengerInput
	orderMessage      contracts.OrderMessageInput
}

// NewService constructor
func NewService() Service {
	st := context.Storage()
	ms := context.Messenger()
	return Service{
		customerMessenger: usecases.NewCustomerMessengerUseCase(st, ms),
		orderMessage:      usecases.NewOrderMessageUseCase(st, ms),
	}
}

// ListenChat listen messenger chat
func (s Service) ListenChat() {
	s.customerMessenger.ListenChat()
}

// OrderMessage send message to phone number
func (s Service) OrderMessage(phone string, message string) {
	s.orderMessage.OrderMessage(phone, message)
}

/*
// OrderMessage ...
func (s Service) OrderMessage(request contracts.OrderMessageRequest, output contracts.OrderMessageOutput) {

}

// OrderMessageResponse ...
type OrderMessageResponse struct {
}

// OnResponse ...
func (r OrderMessageResponse) OnResponse() {

}
*/
