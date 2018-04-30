package download

import (
	"net/http"
	"strings"

	"../../configuration/context"
	"../../contracts/usecases"
	"../../usecases"
)

type service struct {
	takeOrder contracts.TakeOrderInput
	uri       string
}

func newService(uri string) service {
	return service{
		takeOrder: usecases.NewTakeOrderUseCase(context.Storage()),
		uri:       uri,
	}
}

func (s service) downloadOrder(w http.ResponseWriter, r *http.Request) (string, error) {
	link := strings.Replace(r.RequestURI, s.uri, "", 1)
	req := contracts.TakeOrderRequest{Link: link}
	resp := TakeOrderResponse{writer: w}
	s.takeOrder.TakeOrder(req, resp)
	return link, nil
}
