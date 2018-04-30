package upload

import (
	"crypto/md5"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"../../configuration/context"
	"../../contracts/usecases"
	"../../usecases"
	"gopkg.in/cheggaaa/pb.v1"
)

type service struct {
	placeOrder contracts.PlaceOrderInput
}

func newService() service {
	return service{
		placeOrder: usecases.NewPlaceOrderUseCase(context.Storage()),
	}
}

// TODO: try with MultipartReader
func (s service) uploadOrder(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)
	inMD5 := r.FormValue("MD5")
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("can't get upload file %e", err)
		return "", err
	}
	defer file.Close()

	req := contracts.PlaceOrderRequest{}
	req.Filename = handler.Filename
	req.Dir = "./local/incoming/" + inMD5 + "/"
	req.Phone = r.FormValue("phone")
	req.Private = r.FormValue("private") == "private"

	err = os.MkdirAll(req.Dir, os.ModePerm)
	if err != nil {
		log.Printf("can't create directory %s %e", req.Dir, err)
		return "", err
	}

	temp, err := os.OpenFile(req.GetSourceName(), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("can't create file %e", err)
		return "", err
	}

	copyFile(file, temp, int(handler.Size))
	req.MD5 = computeMD5(temp)
	temp.Close()

	resp := PlaceOrderResponse{phone: req.Phone}
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
