package upload

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
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
	inMD5 := r.FormValue("MD5")
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		return "", err
	}
	defer file.Close()

	por := contracts.PlaceOrderRequest{}
	por.Filename = handler.Filename
	por.Dir = "./local/incoming/" + inMD5 + "/"
	por.Phone = r.FormValue("phone")
	por.Private = r.FormValue("private") == "private"

	err = os.MkdirAll(por.Dir, os.ModePerm)
	// TODO error

	temp, err := os.OpenFile(por.GetSourceName(), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("can't create file")
		return "", err
	}
	defer temp.Close()

	copyFile(file, temp, int(handler.Size))
	por.MD5 = computeMD5(temp)

	s.placeOrder.PlaceOrder(por, PlaceOrderResponse{})

	return temp.Name(), nil
}

func copyFile(in multipart.File, out *os.File, dataSize int) {
	bar := pb.New(dataSize).SetUnits(pb.U_BYTES)
	bar.Start()
	reader := bar.NewProxyReader(in)
	io.Copy(out, reader)
	bar.Finish()
}

func computeMD5(file *os.File) []byte {
	var result []byte
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result
	}
	return hash.Sum(result)
}

type PlaceOrderResponse struct {
}

func (r PlaceOrderResponse) OnResponse(orderID int) {
}

func init() {
}
