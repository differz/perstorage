package api

import (
	"fmt"
	"net/http"

	"../../controller"
)

// Controller ...
type Controller struct {
	name string
}

var uri = "/api"

func new() Controller {
	http.HandleFunc(uri, handler)
	return Controller{
		name: uri,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api: Server request, URL %s", r.URL.Path[1:])
}

func init() {
	controller.Register(uri, new())
}
