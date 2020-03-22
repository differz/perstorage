package download

import (
	"log"
	"net/http"

	"perstorage/web"
)

type controller struct {
	name string
}

var uri = "/download/"

func newController() controller {
	http.HandleFunc(uri, handler)
	return controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		srv := newService(uri)
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
