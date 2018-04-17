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

	req := contracts.PlaceOrderRequest{}
	req.Filename = handler.Filename
	req.Dir = "./local/incoming/" + inMD5 + "/"
	req.Phone = r.FormValue("phone")
	req.Private = r.FormValue("private") == "private"

	err = os.MkdirAll(req.Dir, os.ModePerm)
	// TODO error

	temp, err := os.OpenFile(req.GetSourceName(), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("can't create file")
		return "", err
	}
	defer temp.Close()

	copyFile(file, temp, int(handler.Size))
	req.MD5 = computeMD5(temp)

	resp := PlaceOrderResponse{}
	s.placeOrder.PlaceOrder(req, resp)

	return resp.downloadLink, nil
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

// PlaceOrderResponse ...
type PlaceOrderResponse struct {
	downloadLink string
}

// OnResponse ...
func (r PlaceOrderResponse) OnResponse(orderLink string) {
	r.downloadLink = orderLink
}
