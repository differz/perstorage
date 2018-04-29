package web

import (
	"fmt"
	"log"

	"../common"
)

// Controller for web resouces
type Controller interface {
	//Handler(w http.ResponseWriter, r *http.Request)
	//	String() string
}

const component = "controllers"

var controllers = make(map[string]Controller)

// Register new controller
func Register(name string, controller Controller) {
	if controller == nil {
		log.Panicf("controller factory %s does not exist", name)
	}
	_, registered := controllers[name]
	if registered {
		log.Printf("controller %s already registered", name)
	}
	controllers[name] = controller
}

// Get named controller
func Get(name string) (Controller, error) {
	controller, ok := controllers[name]
	if !ok {
		return nil, fmt.Errorf("cnknown controller type: %s", name)
	}
	return controller, nil
}

// String presents view of factory map
func String() string {
	return fmt.Sprint(controllers)
}

// Print ...
func Print() {
	common.ContextUpMessage(component, String())
}
