package usecases

import (
	"log"

	"github.com/differz/perstorage/contracts/usecases"
	"github.com/differz/perstorage/core"
	"github.com/differz/perstorage/storage"
)

// PlaceOrderUseCase object
type PlaceOrderUseCase struct {
	repo storage.Storager
}

// NewPlaceOrderUseCase constructor
func NewPlaceOrderUseCase(repo storage.Storager) PlaceOrderUseCase {
	return PlaceOrderUseCase{
		repo: repo,
	}
}

// PlaceOrder stores all order info and call order link delivery to customer
func (u PlaceOrderUseCase) PlaceOrder(request contracts.PlaceOrderRequest, output contracts.PlaceOrderOutput) {
	var err error

	customerID := request.CustomerID
	if customerID == 0 {
		customerID, err = core.GetCustomerIDByPhone(request.Phone)
		if err != nil {
			log.Printf("can't get customer id by phone %s %e", request.Phone, err)
			return
		}
	}

	customer, ok := u.repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = request.Phone
		u.repo.StoreCustomer(customer)
	}

	item := core.Item{Filename: request.Filename, SourceName: request.GetSourceName(), Category: core.Category(request.CategoryID)}
	item.ID, err = u.repo.StoreItem(item)
	if err != nil {
		log.Printf("can't store item %s %e", item.Filename, err)
		return
	}

	order := core.Order{Customer: customer, Description: request.Description}
	order.Add(item)

	order.ID, err = u.repo.StoreOrder(order)
	if err != nil {
		log.Printf("can't store order for customer %s %e", order.Customer.Name, err)
		return
	}

	output.OnResponse(customer.Phone, order.Link(), order.Description)
}
