package usecases

import (
	"log"

	"../core"
	"../messenger"
	"../storage"
)

// OrderMessageUseCase object
type OrderMessageUseCase struct {
	repo storage.Storager
	msgr messenger.Messenger
}

// NewOrderMessageUseCase constructor
func NewOrderMessageUseCase(repo storage.Storager, msgr messenger.Messenger) OrderMessageUseCase {
	return OrderMessageUseCase{
		repo:        repo,
		msgr:        msgr,
	}
}

// OrderMessage show message to customers messenger by phone number
func (u OrderMessageUseCase) OrderMessage(phone string, message, description string) {
	customerID, err := core.GetCustomerIDByPhone(phone)
	if err != nil {
		log.Printf("can't get customer id by phone %s", phone)
	}

	customer := core.Customer{ID: customerID}
	chatID, ok := u.repo.FindCustomerChatID(customer, u.msgr.Name())
	if !ok {
		log.Printf("no chat id for customer %d", customerID)
	}

	u.msgr.ShowOrder(chatID, message, description)
}

// TODO: request
/*
func (u OrderMessageUseCase) OrderMessage(request contracts.OrderMessageRequest, output contracts.OrderMessageOutput) {
	output.OnResponse(request.Phone, request.Message)
}
*/
