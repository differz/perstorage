package usecases

import (
	"log"

	"../../core"
)

func (r CustomerMessengerResponse) registerMessenger(phone, messenger string, chatID int) {
	customerID, err := core.GetCustomerIDByPhone(phone)
	if err != nil {
		log.Printf("can't get customer id by phone %s %e", phone, err)
		return
	}

	customer, ok := r.repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = phone
		r.repo.StoreCustomer(customer)
	}

	_, ok = r.repo.FindCustomerChatID(customer, messenger)
	if !ok {
		r.repo.StoreCustomerMessenger(customer, messenger, chatID)
	}
}
