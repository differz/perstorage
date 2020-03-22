package api

import (
	"github.com/differz/perstorage/configuration/context"
	"github.com/differz/perstorage/contracts/usecases"
	"github.com/differz/perstorage/usecases"
)

type service struct {
	older contracts.PurgeOrdersInput
	days  int
}

func newService(uri string) service {
	return service{
		older: usecases.NewPurgeOrdersOlderUseCase(context.Storage(), 30),
		days:  30,
	}
}

func (s service) purgeOrders() error {
	s.older.PurgeOrders()
	return nil
}
