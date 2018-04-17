package usecases

import (
	"fmt"
	"strconv"
	"strings"

	"../configuration"
	"../contracts/usecases"
	"../core"
)

type PlaceOrderUseCase struct {
	//
	subject     string
	description string
}

func NewPlaceOrderUseCase() PlaceOrderUseCase {
	return PlaceOrderUseCase{
		description: "new",
	}
}

func (u PlaceOrderUseCase) PlaceOrder(request contracts.PlaceOrderRequest, output contracts.PlaceOrderOutput) {
	repo := configuration.Get().Storage

	customerID, err := strconv.Atoi(strings.Replace(request.Phone, "+", "", 1))
	if err != nil {
		// TODO error
	}

	customer, ok := repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = request.Phone
		repo.StoreCustomer(customer)
	}

	fmt.Println(customer)

	item := core.Item{Filename: request.Filename, SourceName: request.GetSourceName()}
	item.ID, err = repo.StoreItem(item)
	if err != nil {
		return
	}

	order := core.Order{Customer: customer}
	order.Add(item)

	order.ID, err = repo.StoreOrder(order)
	if err != nil {
		return
	}

	output.OnResponse(order.ID)
}
