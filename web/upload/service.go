package upload

import (
	"io"
	"net/http"
	"os"

	"../../contracts/usecases"
	"../../storage"
	"../../usecases"
)

// Service ...
type Service struct {
	storage    storage.Storager
	placeOrder contracts.PlaceOrderInput
}

// NewService constructor
func NewService(storage storage.Storager) Service {
	return Service{
		storage:    storage,
		placeOrder: usecases.NewPlaceOrderUseCase(),
	}
}

func (s Service) uploadFile(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		return "", err
	}
	defer file.Close()
	f, err := os.OpenFile("./local/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	por := contracts.PlaceOrderRequest{}
	por.Filename = handler.Filename
	por.Subject = "subject"

	s.placeOrder.PlaceOrder(por, PlaceOrderResponse{})

	return f.Name(), nil
}

type PlaceOrderResponse struct {
}

func (r PlaceOrderResponse) OnResponse(orderID int64) {
}

func init() {
}
