package upload

import (
	"io"
	"net/http"
	"os"

	"../../contracts/usecases"
	"../../usecases"
	"gopkg.in/cheggaaa/pb.v1"
)

// Service ...
type Service struct {
	placeOrder contracts.PlaceOrderInput
}

// NewService constructor
func NewService() Service {
	return Service{
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

	temp, err := os.OpenFile("./local/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer temp.Close()

	dataSize := int(handler.Size)
	bar := pb.New(dataSize).SetUnits(pb.U_BYTES)
	bar.Start()
	reader := bar.NewProxyReader(file)
	io.Copy(temp, reader)
	bar.Finish()

	por := contracts.PlaceOrderRequest{}
	por.Filename = handler.Filename
	por.Phone = r.FormValue("phone")
	por.Private = r.FormValue("private") == "private"

	s.placeOrder.PlaceOrder(por, PlaceOrderResponse{})

	return temp.Name(), nil
}

type PlaceOrderResponse struct {
}

func (r PlaceOrderResponse) OnResponse(orderID int) {
}

func init() {
}
