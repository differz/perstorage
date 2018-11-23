package file

import (
	"log"
	"strings"
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
	if err != nil {
		log.Fatal(err)
	}
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

	order, customerID, ok := s.getOrderHeadByID(id)
	if !ok {
		return order, ok
	}

	order, ok = s.fillOrderItems(order)
	if !ok {
		return order, ok
	}

	order.Customer, ok = s.findCustomerByID(customerID)
	return order, ok
}

func (s Storage) getOrderHeadByID(id int) (core.Order, int, bool) {
	sql := "SELECT id, description, size, category, order_date, customer_id FROM orders WHERE id = ?"
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
	customerID := 0
	ok := false
	if rows.Next() {
		err = rows.Scan(&order.ID, &order.Description, &order.Size, &order.Category, &order.Date, &customerID)
		if err != nil {
			log.Fatal(err)
		} else {
			ok = true
		}
	}
	return order, customerID, ok
}

func (s Storage) fillOrderItems(order core.Order) (core.Order, bool) {
	items, ok := s.GetOrderedItems(order)
	if ok {
		for _, item := range items {
			order.Add(item)
		}
	}
	return order, ok
}

func (s Storage) findCustomerByID(id int) (core.Customer, bool) {
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
		} else {
			ok = true
		}
	}
	return customer, ok
}

// FindOrderByLink get bucket from storage by link
func (s Storage) FindOrderByLink(link string) (core.Order, bool) {
	orderID, ok := s.findOrderIDByLink(link)
	order := core.Order{}
	if !ok {
		return order, false
	}
	return s.FindOrderByID(orderID)
}

func (s Storage) findOrderIDByLink(link string) (int, bool) {
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
		return 0, false
	}
	return order.ID, true
}

// GetOrders takes oreders from db
func (s Storage) GetOrders(strategy func() string) ([]core.Order, error) {
	sql := "SELECT" +
		"   o.id, o.description, o.size, o.category, o.order_date," +
		"   c.id, c.name, c.phone" +
		" FROM orders AS o" +
		" LEFT JOIN customers AS c" +
		"	ON o.customer_id = c.id" +
		" WHERE true"
	where := strategy()
	if where != "" {
		sql = strings.Replace(sql, "true", where, 1)
	}

	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var heads []core.Order
	for rows.Next() {
		customer := core.Customer{}
		order := core.Order{}
		err = rows.Scan(&order.ID, &order.Description, &order.Size, &order.Category, &order.Date,
			&customer.ID, &customer.Name, &customer.Phone)
		if err != nil {
			log.Fatal(err)
		}
		order.Customer = customer
		heads = append(heads, order)
	}

	var orders []core.Order
	for _, order := range heads {
		order, _ = s.fillOrderItems(order)
		orders = append(orders, order)
	}
	return orders, nil
}

// GetOrderedItems takes all ordered items by selected order
func (s Storage) GetOrderedItems(order core.Order) ([]core.Item, bool) {
	sql := "SELECT" +
		"  i.id," +
		"  i.name," +
		"  i.filename," +
		"  i.path," +
		"  i.size," +
		"  i.category," +
		"  i.available" +
		" FROM ordered_items AS oi" +
		" LEFT JOIN items AS i" +
		"	ON oi.item_id = i.id" +
		" WHERE order_id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(order.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []core.Item
	ok := false
	for rows.Next() {
		item := core.Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Filename, &item.SourceName, &item.Size, &item.Category, &item.Available)
		if err != nil {
			log.Print(err)
		} else {
			ok = true
			items = append(items, item)
		}
	}
	return items, ok
}

// DeleteOrder remove order from storage
func (s Storage) DeleteOrder(order core.Order) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if !s.deleteOrder(order) {
		return false
	}
	if !s.deleteOrderLink(order) {
		return false
	}
	if !s.deleteOrderedItems(order) {
		return false
	}
	return true
}

func (s Storage) deleteOrder(order core.Order) bool {
	sql := "DELETE FROM orders WHERE id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table orders ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID)
	if err != nil {
		log.Fatal("can't delete order from db ", err)
	}
	return true
}

func (s Storage) deleteOrderLink(order core.Order) bool {
	sql := "DELETE FROM order_links WHERE order_id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table order_links ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID)
	if err != nil {
		log.Fatal("can't delete order_links from db ", err)
	}
	return true
}

func (s Storage) deleteOrderedItems(order core.Order) bool {
	sql := "DELETE FROM ordered_items WHERE order_id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table ordered_items ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID)
	if err != nil {
		log.Fatal("can't delete ordered_items from db ", err)
	}
	return true
}
