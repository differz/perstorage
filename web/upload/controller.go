package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"../../web"
)

// Controller ...
type Controller struct {
	name string
}

var uri = "/upload"

func new() Controller {
	http.HandleFunc(uri, handler)
	return Controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "/upload: Server request, URL %s", r.URL.Path[1:])

	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./local/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

	http.Redirect(w, r, "/", 301)
}

func init() {
	web.Register(uri, new())
}
