package usecases

import (
	"../configuration"
	"../contracts/usecases"
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
	link := request.Link
	repo := configuration.GetStorage()
	order, ok := repo.FindOrderByLink(link)
	if ok {
		for _, item := range order.Items {
			output.OnResponse(item.SourceName, item.Filename, item.Size)
		}
	}
}
