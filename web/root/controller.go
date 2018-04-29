package root

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../../web"
)

// Controller for root
type Controller struct {
	name string
}

var uri = "/"

// fmt.Fprintf(w, "default: Server request, URL %s", r.URL.Path[1:])
func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Body)
}

func new() Controller {
	c := Controller{
		name: uri,
	}
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	http.Handle(uri, web.Adapt(c, web.Notify(logger)))
	return c
}

func init() {
	web.Register("root", new())
}
