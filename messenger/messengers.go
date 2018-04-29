package messenger

import (
	"fmt"
	"log"

	"../common"
	"../contracts/messengers"
)

// Messenger consist of all messengers contracts that each should implement
type Messenger interface {
	messengers.ListenChatInput
	messengers.OrderPostInput
	Init(args ...string) error
	Available() bool
}

const component = "messenger"

var ms = make(map[string]Messenger)

// Register messenger into factory by name
func Register(name string, messenger Messenger) {
	if messenger == nil {
		log.Panicf("messenger factory %s does not exist", name)
	}
	_, registered := ms[name]
	if registered {
		log.Printf("messenger %s already registered", name)
	}
	ms[name] = messenger
}

// Get messenger by name from factory
func Get(name string, args ...string) (Messenger, error) {
	messenger, ok := ms[name]
	if !ok {
		return nil, fmt.Errorf("unknown messenger type: %s", name)
	}
	messenger.Init(args...)
	return messenger, nil
}

// String presents view of factory map
func String() string {
	return fmt.Sprint(ms)
}

// Print view of factory map to console
func Print() {
	common.ContextUpMessage(component, String())
}
