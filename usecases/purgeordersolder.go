package usecases

import (
	"fmt"
	"log"

	"github.com/differz/perstorage/configuration"
	"github.com/differz/perstorage/storage"
)

// PurgeOrdersOlderUseCase object
type PurgeOrdersOlderUseCase struct {
	repo        storage.Storager
	days        int
	description string
}

// NewPurgeOrdersOlderUseCase constructor
func NewPurgeOrdersOlderUseCase(repo storage.Storager, days int) PurgeOrdersOlderUseCase {
	return PurgeOrdersOlderUseCase{
		repo:        repo,
		days:        days,
		description: "new",
	}
}

// PurgeOrders delete orders by clause
func (u PurgeOrdersOlderUseCase) PurgeOrders() {
	orders, err := u.repo.GetOrders(strategyOldOrders)
	if err != nil {
		log.Fatal(err)
	}

	for _, order := range orders {
		fmt.Printf("| delete | %s", order.String())
		for _, item := range order.Items {
			u.repo.DeleteItem(item)
		}
		u.repo.DeleteOrder(order)
		break
	}
}

func strategyOldOrders() string {
	return "order_date < '" + configuration.PurgeDate() + "'"
}
