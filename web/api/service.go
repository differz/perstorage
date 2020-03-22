package api

import (
	"perstorage/configuration/context"
	"perstorage/contracts/usecases"
	"perstorage/usecases"
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
