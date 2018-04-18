package usecases

import (
	"../contracts/usecases"
)

// OrderMessageUseCase ...
type OrderMessageUseCase struct {
	//
	subject     string
	description string
}

// NewOrderMessageUseCase ...
func NewOrderMessageUseCase() OrderMessageUseCase {
	return OrderMessageUseCase{
		description: "new",
	}
}

// OrderMessage ...
func (u OrderMessageUseCase) OrderMessage(request contracts.OrderMessageRequest, output contracts.OrderMessageOutput) {

	output.OnResponse("order.Link()")
}
