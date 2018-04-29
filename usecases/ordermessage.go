package usecases

import (
	"fmt"
	"strconv"
	"strings"

	"../configuration/context"
	"../core"
)

// OrderMessageUseCase ...
type OrderMessageUseCase struct {
	//
	subject     string
	description string
}

// NewOrderMessageUseCase ...
func NewOrderMessageUseCase() OrderMessageUseCase {
	return OrderMessageUseCase{
		description: "new",
	}
}

// OrderMessage ...
func (u OrderMessageUseCase) OrderMessage(phone string, message string) {
	// TODO: @Inject
	repo := context.Storage()
	msgr := context.Messenger()

	// TODO: find customer by phone
	customerID, _ := strconv.Atoi(strings.Replace(phone, "+", "", 1))

	customer := core.Customer{ID: customerID}
	chatID, ok := repo.FindCustomerChatID(customer, "telegram") //TODO: @Inject
	if !ok {
		fmt.Printf("no chat id for customer %d", customerID)
	}

	msgr.ShowOrder(chatID, message)

}

// TODO: request
/*
func (u OrderMessageUseCase) OrderMessage(request contracts.OrderMessageRequest, output contracts.OrderMessageOutput) {
	output.OnResponse(request.Phone, request.Message)
}
*/
