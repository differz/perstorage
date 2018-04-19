package messenger

import (
	"fmt"
	"log"

	"../contracts/messengers"
)

// Storager ...
type Messenger interface {
	messengers.OrderPost
}

var ms = make(map[string]Messenger)

// Register ...
func Register(name string, messenger Messenger) {
	if messenger == nil {
		log.Panicf("Messenger factory %s does not exist", name)
	}
	_, registered := ms[name]
	if registered {
		log.Printf("Messenger %s already registered", name)
	}
	ms[name] = messenger
}

// Get ...
func Get(name string, args ...string) (Messenger, error) {
	messenger, ok := ms[name]
	if !ok {
		return nil, fmt.Errorf("Unknown messenger type: %s", name)
	}
	return messenger, nil
}

// Print ...
func Print() {
	fmt.Println(ms)
}
