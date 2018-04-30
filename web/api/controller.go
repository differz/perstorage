package api

import (
	"fmt"
	"net/http"

	"../../web"
)

type controller struct {
	name string
}

var uri = "/api"

func newController() controller {
	http.HandleFunc(uri, handler)
	return controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api: Server request, URL %s", r.URL.Path[1:])
}

func init() {
	web.Register(uri, newController())
}
