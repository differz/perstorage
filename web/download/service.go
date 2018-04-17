package download

import (
	"io"
	"net/http"
	"os"
	"strconv"
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

// TakeOrderResponse ...
type TakeOrderResponse struct {
	writer http.ResponseWriter
}

// OnResponse ...
func (r TakeOrderResponse) OnResponse(sourcename, filename string, size int) {
	file, err := os.Open(sourcename)
	if err != nil {
		http.Error(r.writer, "File not found.", 404)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileType := http.DetectContentType(fileHeader)

	length := strconv.FormatInt(int64(size), 10)

	header := r.writer.Header()
	header.Set("Content-Disposition", "attachment; filename="+filename)
	header.Set("Content-Type", fileType)
	header.Set("Content-Length", length)

	file.Seek(0, 0)
	io.Copy(r.writer, file)
}
