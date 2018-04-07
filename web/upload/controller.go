package upload

import (
	"fmt"
	"net/http"

	"../../web"
)

// Controller ...
type Controller struct {
	name string
}

var uri = "/upload"

func newController() Controller {
	http.HandleFunc(uri, handler)
	return Controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "/upload: Server request, URL %s", r.URL.Path[1:])
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		//	repo, _ := storage.Get("file", 1)
		srv := NewService()

		name, err := srv.uploadFile(r)
		if err == nil {
			fmt.Println(name)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func init() {
	web.Register(uri, newController())
}
