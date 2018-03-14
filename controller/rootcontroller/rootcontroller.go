package rootcontroller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	co "../../controller"
)

// Controller ...
type Controller struct {
	name string
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "default: Server request, URL %s", r.URL.Path[1:])
}

func new() Controller {
	c := Controller{
		name: "/",
	}

	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	http.Handle("/", co.Adapt(c, co.Notify(logger)))

	return c
}

func init() {
	co.Register("root", new())
}
