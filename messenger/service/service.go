package messengers

import (
	"../../configuration"
	"../../contracts/usecases"
	"../../usecases"
)

// Service ...
type Service struct {
	customerMessenger contracts.CustomerMessengerInput
	orderMessage      contracts.OrderMessageInput
}

// NewService constructor
func NewService() Service {
	return Service{
		customerMessenger: usecases.NewCustomerMessengerUseCase(configuration.GetStorage(), configuration.GetMessenger()),
		orderMessage:      usecases.NewOrderMessageUseCase(),
	}
}

func (s Service) ListenChat() {
	s.customerMessenger.ListenChat()
}

// OrderMessage ...
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
