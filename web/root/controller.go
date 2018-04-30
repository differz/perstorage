package root

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../../web"
)

type controller struct {
	name string
}

var uri = "/"

func (c controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Body)
}

func newController() controller {
	c := controller{
		name: uri,
	}
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	http.Handle(uri, web.Adapt(c, web.Notify(logger)))
	return c
}

func init() {
	web.Register("root", newController())
}
