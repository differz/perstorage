package configuration

import (
	"encoding/json"
	"log"
	"os"
	"strings"
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
	cfgFile   = "perstorage.json"
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

// ServerPort current server port
func ServerPort() string {
	port := ":80"
	addr := ServerAddress()
	idx := strings.LastIndex(addr, ":")
	if idx > 10 {
		port = addr[idx:]
	}
	return port
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
