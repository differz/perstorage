package download

import (
	"net/http"

	"../../contracts/usecases"
	"../../usecases"
)

// Service ...
type Service struct {
	takeOrder contracts.TakeOrderInput
}

// NewService constructor
func NewService() Service {
	return Service{
		takeOrder: usecases.NewTakeOrderUseCase(),
	}
}

func (s Service) downloadFile(r *http.Request) (string, error) {

	tor := contracts.TakeOrderRequest{}

	s.takeOrder.TakeOrder(tor, TakeOrderResponse{})

	return "", nil
}

type TakeOrderResponse struct {
}

func (r TakeOrderResponse) OnResponse(orderID int) {
}

func init() {
}
