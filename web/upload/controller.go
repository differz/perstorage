package upload

import (
	"net/http"

	"../../web"
	"log"
)

// Controller object
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
	if r.Method == "POST" {
		srv := NewService()
		name, err := srv.uploadOrder(r)
		if err == nil {
			log.Printf("uploaded order %s", name)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func init() {
	web.Register(uri, newController())
}
