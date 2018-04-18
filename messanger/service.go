package messanger

import (
	"../contracts/usecases"
	"../usecases"
)

// Service ...
type Service struct {
	orderMessage contracts.OrderMessageInput
}

// NewService constructor
func NewService() Service {
	return Service{
		orderMessage: usecases.NewOrderMessageUseCase(),
	}
}

func (s Service) OrderMessage(request contracts.OrderMessageRequest, output contracts.OrderMessageOutput) {

}

// OrderMessageResponse ...
type OrderMessageResponse struct {
}

// OnResponse ...
func (r OrderMessageResponse) OnResponse() {

}
