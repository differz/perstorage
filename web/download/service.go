package download

import (
	"net/http"
	"strings"

	"../../contracts/usecases"
	"../../usecases"
)

// Service ...
type Service struct {
	takeOrder contracts.TakeOrderInput
	uri       string
}

// NewService constructor
func NewService(uri string) Service {
	return Service{
		takeOrder: usecases.NewTakeOrderUseCase(),
		uri:       uri,
	}
}

func (s Service) downloadFile(w http.ResponseWriter, r *http.Request) (string, error) {
	link := strings.Replace(r.RequestURI, s.uri, "", 1)
	req := contracts.TakeOrderRequest{Link: link}
	resp := TakeOrderResponse{writer: w}
	s.takeOrder.TakeOrder(req, resp)

	return "", nil
}
