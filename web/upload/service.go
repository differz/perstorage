package service

import (
	"../storage"
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

// Save file to storage
func (s *Service) Save() {
	s.storage.Save()
}

// Get file from storage
func (s *Service) Get() {
	s.storage.Get()
}

func init() {

}
