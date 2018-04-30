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

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Body)
}

func newController() Controller {
	c := Controller{
		name: uri,
	}
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	http.Handle(uri, web.Adapt(c, web.Notify(logger)))
	return c
}

func init() {
	web.Register("root", newController())
}
