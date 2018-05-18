package file

import (
	"log"

	"../../core"
)

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(customer core.Customer) (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	sql := "INSERT INTO customers(id, name, phone) VALUES(?, ?, ?)"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(customer.ID, customer.Name, customer.Phone)
	if err != nil {
		log.Fatal(err)
	}

	if customer.IsNew() {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		}
		customer.ID = int(id)
	}
	return customer.ID, err
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	sql := "SELECT id, name, phone FROM customers WHERE id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	customer := core.Customer{}
	ok := false
	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone)
		if err != nil {
			log.Fatal(err)
		}
		ok = true
	}
	return customer, ok
}
