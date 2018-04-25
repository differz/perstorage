package usecases

import (
	"../contracts/usecases"
	"../storage"
)

// TakeOrderUseCase ...
type TakeOrderUseCase struct {
	repo storage.Storager
	//
	subject     string
	description string
}

// NewTakeOrderUseCase ...
func NewTakeOrderUseCase(repo storage.Storager) TakeOrderUseCase {
	return TakeOrderUseCase{
		repo:        repo,
		description: "new",
	}
}

// TakeOrder ...
func (u TakeOrderUseCase) TakeOrder(request contracts.TakeOrderRequest, output contracts.TakeOrderOutput) {
	link := request.Link
	//repo := configuration.GetStorage()
	order, ok := u.repo.FindOrderByLink(link)
	if ok {
		for _, item := range order.Items {
			output.OnResponse(item.SourceName, item.Filename, item.Size)
		}
	}
}
