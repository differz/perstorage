package api

import (
	"fmt"
	"log"
	"net/http"

	"perstorage/web"
)

type controller struct {
	name string
}

var uri = "/api/"

func newController() controller {
	http.HandleFunc(uri, handler)
	return controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/: Server request, URL %s", r.URL.Path[1:])
	if r.URL.Path[1:] == "api/purge" {
		srv := newService(uri)
		err := srv.purgeOrders()
		if err == nil {
			log.Println("purge orders")
		}
	}
}

func init() {
	web.Register(uri, newController())
}
