package usecases

import (
	"../configuration"
	"../contracts/usecases"
	"../core"
)

type PlaceOrderUseCase struct {

	//
	projectId   int64
	subject     string
	description string
}

func NewPlaceOrderUseCase() PlaceOrderUseCase {
	return PlaceOrderUseCase{
		projectId: 1,
	}
}

func (u PlaceOrderUseCase) PlaceOrder(request contracts.PlaceOrderRequest, output contracts.PlaceOrderOutput) {
	filename := request.Filename

	orderID := int64(1)

	repo := configuration.Get().Storage

	repo.StoreItem(core.Item{Filename: filename})

	output.OnResponse(orderID)
}
