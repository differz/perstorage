package download

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

// TakeOrderResponse response data
type TakeOrderResponse struct {
	writer http.ResponseWriter
}

// OnResponse write file and fill header to writer handler
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
