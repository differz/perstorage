package upload

import (
	"log"
	"net/http"

	"../../web"
)

type controller struct {
	name string
}

var uri = "/upload"

func newController() controller {
	http.HandleFunc(uri, handler)
	return controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		srv := newService()
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
