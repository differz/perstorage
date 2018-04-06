package controller

import (
	"fmt"
	"log"
)

// Controller for web resouces
type Controller interface {
	//Handler(w http.ResponseWriter, r *http.Request)
	//	String() string
}

var controllers = make(map[string]Controller)

// Register new controller
func Register(name string, controller Controller) {
	if controller == nil {
		log.Panicf("Controller factory %s does not exist", name)
	}
	_, registered := controllers[name]
	if registered {
		log.Printf("Controller %s already registered", name)
	}
	controllers[name] = controller
}

// Get named controller
func Get(name string, x int) (Controller, error) {
	controller, ok := controllers[name]
	if !ok {
		return nil, fmt.Errorf("Unknown controller type: %s", name)
	}
	return controller, nil
}

//func PostConstruct()

// Print ...
func Print() {
	fmt.Println(controllers)
}
