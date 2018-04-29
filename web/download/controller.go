package download

import (
	"fmt"
	"net/http"

	"../../web"
)

// Controller object
type Controller struct {
	name string
}

var uri = "/download/"

func newController() Controller {
	http.HandleFunc(uri, handler)
	return Controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		srv := NewService(uri)
		name, err := srv.downloadFile(w, r)
		if err == nil {
			fmt.Println(name)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func init() {
	web.Register(uri, newController())
}
