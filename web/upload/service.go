package upload

import (
	"../../storage"
)

type Service struct {
	storage storage.Storager
	hash    []byte
}

// New constructor
func New(storage storage.Storager) *Service {
	return &Service{
		storage: storage,
	}
}

func init() {

}
