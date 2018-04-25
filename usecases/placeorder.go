package usecases

import (
	"fmt"
	"strconv"
	"strings"

	"../contracts/usecases"
	"../core"
	"../storage"
)

// PlaceOrderUseCase ...
type PlaceOrderUseCase struct {
	repo storage.Storager
	//
	subject     string
	description string
}

// NewPlaceOrderUseCase ...
func NewPlaceOrderUseCase(repo storage.Storager) PlaceOrderUseCase {
	return PlaceOrderUseCase{
		repo:        repo,
		description: "new",
	}
}

// PlaceOrder ...
func (u PlaceOrderUseCase) PlaceOrder(request contracts.PlaceOrderRequest, output contracts.PlaceOrderOutput) {

	customerID, err := strconv.Atoi(strings.Replace(request.Phone, "+", "", 1))
	if err != nil {
		// TODO error
	}

	customer, ok := u.repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = request.Phone
		u.repo.StoreCustomer(customer)
	}

	fmt.Println(customer)

	item := core.Item{Filename: request.Filename, SourceName: request.GetSourceName()}
	item.ID, err = u.repo.StoreItem(item)
	if err != nil {
		return
	}

	order := core.Order{Customer: customer}
	order.Add(item)

	order.ID, err = u.repo.StoreOrder(order)
	if err != nil {
		return
	}

	output.OnResponse(customer.Phone, order.Link())
}
