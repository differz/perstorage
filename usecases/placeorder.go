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
	filename := request.Filename

	repo := configuration.Get().Storage

	customerID, err := strconv.Atoi(strings.Replace(request.Phone, "+380", "", 1))
	if err != nil {

	}

	customer, ok := repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = request.Phone
		repo.StoreCustomer(customer)
	}

	fmt.Println(customer)

	item := core.Item{Filename: filename}
	repo.StoreItem(item)

	order := core.Order{Customer: customer}
	order.Add(item)

	repo.StoreOrder(order)

	output.OnResponse(order.ID)
}
