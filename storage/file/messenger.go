package file

import (
	"log"

	"../../core"
)

// StoreCustomerMessenger save chat id to storage by customer
func (s Storage) StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error {
	mutex.Lock()
	defer mutex.Unlock()

	sql := "INSERT INTO customer_messengers(customer_id, messenger, chat_id) VALUES(?, ?, ?)"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.ID, messengerName, chatID)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// FindCustomerChatID get customer's chat id
func (s Storage) FindCustomerChatID(customer core.Customer, messengerName string) (int, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	sql := "SELECT chat_id FROM customer_messengers WHERE customer_id = ? AND messenger = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(customer.ID, messengerName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	chatID := 0
	ok := false
	if rows.Next() {
		err = rows.Scan(&chatID)
		if err != nil {
			log.Fatal(err)
		}
		ok = true
	}
	return chatID, ok
}

// IsRegisteredChatID get customer's id if chat already registered
func (s Storage) IsRegisteredChatID(chatID int, messengerName string) (int, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	sql := "SELECT customer_id FROM customer_messengers WHERE chat_id = ? AND messenger = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(chatID, messengerName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	customerID := 0
	ok := false
	if rows.Next() {
		err = rows.Scan(&customerID)
		if err != nil {
			log.Fatal(err)
		}
		ok = true
	}
	return customerID, ok
}
