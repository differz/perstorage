package messengers

import (
	"github.com/differz/perstorage/configuration/context"
	"github.com/differz/perstorage/contracts/usecases"
	uc "github.com/differz/perstorage/usecases"
	ucm "github.com/differz/perstorage/usecases/messenger"
)

// Service object
type Service struct {
	customerMessenger contracts.CustomerMessengerInput
	orderMessage      contracts.OrderMessageInput
}

// NewService constructor
func NewService() Service {
	repo := context.Storage()
	msgr := context.Messenger()
	return Service{
		customerMessenger: ucm.NewCustomerMessengerUseCase(repo, msgr),
		orderMessage:      uc.NewOrderMessageUseCase(repo, msgr),
	}
}

// ListenChat listen messenger chat
func (s Service) ListenChat() {
	s.customerMessenger.ListenChat()
}

// OrderMessage send message to phone number
func (s Service) OrderMessage(phone string, message, description string) {
	s.orderMessage.OrderMessage(phone, message, description)
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
