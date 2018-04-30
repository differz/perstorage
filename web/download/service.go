package download

import (
	"net/http"
	"strings"

	"../../configuration/context"
	"../../contracts/usecases"
	"../../usecases"
)

// Service object
type Service struct {
	takeOrder contracts.TakeOrderInput
	uri       string
}

// NewService constructor
func NewService(uri string) Service {
	return Service{
		takeOrder: usecases.NewTakeOrderUseCase(context.Storage()),
		uri:       uri,
	}
}

func (s Service) downloadOrder(w http.ResponseWriter, r *http.Request) (string, error) {
	link := strings.Replace(r.RequestURI, s.uri, "", 1)
	req := contracts.TakeOrderRequest{Link: link}
	resp := TakeOrderResponse{writer: w}
	s.takeOrder.TakeOrder(req, resp)
	return link, nil
}
