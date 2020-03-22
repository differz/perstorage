package root

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"perstorage/configuration"
	"perstorage/web"
)

type controller struct {
	name string
}

var uri = "/"

func (c controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := strings.Replace(Body, "hostname", configuration.ServerAddress(), 1)
	fmt.Fprintf(w, body)
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
