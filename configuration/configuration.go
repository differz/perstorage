package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"../common"
)

type config struct {
	ConfigName    string
	StorageName   string
	StorageArgs   string
	MessengerName string
	MessengerKey  string
}

const (
	component = "configuration"
	cfgFile   = "config.json"
)

var (
	cfg  *config
	once sync.Once
)

func Get() *config {
	once.Do(func() {
		cfg = &config{}
		cfg.read()
	})
	return cfg
}

func (conf *config) read() {
	file, err := os.Open(cfgFile)
	if err != nil {
		msg := fmt.Sprintf("can't read file %s", cfgFile)
		common.ContextUpMessage(component, msg)
		log.Fatal(msg)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		msg := fmt.Sprintf("error parsing file %s, %e ", cfgFile, err)
		common.ContextUpMessage(component, msg)
		log.Fatal(msg)
	}
	common.ContextUpMessage(component, fmt.Sprint(cfg))
}
