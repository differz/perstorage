package file

import (
	"log"
	"time"

	"../../core"
)

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(order core.Order) (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	tx, err := s.connection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO orders(customer_id, description, order_date) VALUES(?, ?, ?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	now.Format("2006-01-02T15:04:05.999999999")

	res, err := stmt.Exec(order.Customer.ID, order.Description, now)
	if err != nil {
		log.Fatal(err)
	}

	if order.IsNew() {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		}
		order.ID = int(id)
	}

	sql = "INSERT INTO order_links(order_id, link) VALUES(?, ?)"
	stmt, err = tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(order.ID, order.Link())
	if err != nil {
		log.Fatal(err)
	}

	sql = "INSERT INTO ordered_items(order_id, item_id) VALUES(?, ?)"
	stmt, err = tx.Prepare(sql)
	for _, item := range order.Items {
		_, err := stmt.Exec(order.ID, item.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	return order.ID, err
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	sql := "SELECT id, customer_id FROM orders WHERE id = ?"
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

	order := core.Order{}
	customer := core.Customer{}
	if rows.Next() {
		err = rows.Scan(&order.ID, &customer.ID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return order, false
	}

	sql = "SELECT id, name, phone FROM customers WHERE id = ?"
	stmt, err = s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err = stmt.Query(customer.ID)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone)
		if err != nil {
			log.Fatal(err)
		}
		order.Customer = customer
	} else {
		return order, false
	}

	sql = "SELECT" +
		"  i.id," +
		"  i.name," +
		"  i.filename," +
		"  i.path," +
		"  i.size," +
		"  i.available" +
		" FROM ordered_items AS oi" +
		" LEFT JOIN items AS i" +
		"	ON oi.item_id = i.id" +
		" WHERE order_id = ?"
	stmt, err = s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err = stmt.Query(order.ID)
	if err != nil {
		log.Fatal(err)
	}
	ok := false
	for rows.Next() {
		item := core.Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Filename, &item.SourceName, &item.Size, &item.Available)
		if err != nil {
			log.Fatal(err)
		}
		ok = true
		order.Add(item)
	}
	return order, ok
}

// FindOrderByLink get bucket from storage by link
func (s Storage) FindOrderByLink(link string) (core.Order, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	sql := "SELECT order_id FROM order_links WHERE link = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(link)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	order := core.Order{}
	if rows.Next() {
		err = rows.Scan(&order.ID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return order, false
	}
	return s.FindOrderByID(order.ID)
}
