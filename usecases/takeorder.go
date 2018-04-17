package usecases

import (
	"fmt"

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
			fmt.Println(item.SourceName)
		}
	}

	output.OnResponse(order.ID)
}
