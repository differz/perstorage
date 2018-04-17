package usecases

import (
	"crypto/sha256"
	"encoding/hex"
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
	repo := configuration.GetStorage()

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

	orderLink := generateLink(order.ID, customer.ID)
	output.OnResponse(orderLink)
}

func generateLink(orderID, customerID int) string {
	key := "order#" + string(orderID) + ", customer:" + string(customerID)
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)
	return hashHex
}
