package configuration

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Config application parameters
type Config struct {
	ConfigName    string
	ServerAddress string
	StorageName   string
	StorageArgs   string
	MessengerName string
	MessengerKey  string
}

const (
	Component = "configuration"
	cfgFile   = "config.json"
)

var (
	cfg  *Config
	once sync.Once
)

// Get configuration from file
func Get() *Config {
	once.Do(func() {
		cfg = &Config{}
		cfg.read()
	})
	return cfg
}

// Name of configuration
func Name() string {
	return Get().ConfigName
}

// ServerAddress current server ip address with port
func ServerAddress() string {
	return Get().ServerAddress
}

func (conf *Config) read() {
	file, err := os.Open(cfgFile)
	if err != nil {
		log.Fatalf("can't read file %s", cfgFile)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		log.Fatalf("error parsing file %s, %e ", cfgFile, err)
	}
}
