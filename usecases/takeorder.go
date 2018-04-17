package usecases

import (
	"../contracts/usecases"
	"../core"
)

// TakeOrderUseCase ...
type TakeOrderUseCase struct {
	//
	subject     string
	description string
}

// NewTakeOrderUseCase ...
func NewTakeOrderUseCase() TakeOrderUseCase {
	return TakeOrderUseCase{
		description: "new",
	}
}

// TakeOrder ...
func (u TakeOrderUseCase) TakeOrder(request contracts.TakeOrderRequest, output contracts.TakeOrderOutput) {
	//filename := request.Filename
	//repo := configuration.Get().Storage

	order := core.Order{}

	output.OnResponse(order.ID)
}
