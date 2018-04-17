package usecases

import (
	"../contracts/usecases"
	"../core"
)

type TakeOrderUseCase struct {
	//
	subject     string
	description string
}

func NewTakeOrderUseCase() TakeOrderUseCase {
	return TakeOrderUseCase{
		description: "new",
	}
}

func (u TakeOrderUseCase) TakeOrder(request contracts.TakeOrderRequest, output contracts.TakeOrderOutput) {
	//filename := request.Filename
	//repo := configuration.Get().Storage

	order := core.Order{}

	output.OnResponse(order.ID)
}
