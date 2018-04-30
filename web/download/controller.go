package download

import (
	"log"
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
		name, err := srv.downloadOrder(w, r)
		if err == nil {
			log.Printf("downloaded order %s", name)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func init() {
	web.Register(uri, newController())
}
