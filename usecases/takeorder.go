package usecases

import (
	"perstorage/contracts/usecases"
	"perstorage/storage"
)

// TakeOrderUseCase object
type TakeOrderUseCase struct {
	repo storage.Storager
	//
	subject     string
	description string
}

// NewTakeOrderUseCase constructor
func NewTakeOrderUseCase(repo storage.Storager) TakeOrderUseCase {
	return TakeOrderUseCase{
		repo:        repo,
		description: "new",
	}
}

// TakeOrder takes order by link and send to response all ordered items
func (u TakeOrderUseCase) TakeOrder(request contracts.TakeOrderRequest, output contracts.TakeOrderOutput) {
	order, ok := u.repo.FindOrderByLink(request.Link)
	if ok {
		for _, item := range order.Items {
			output.OnResponse(item.SourceName, item.Filename, item.Size)
		}
	}
}
