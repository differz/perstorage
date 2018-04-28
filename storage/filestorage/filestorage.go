package filestorage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	"github.com/satori/go.uuid"
	// sqlite
	_ "github.com/mattes/migrate/source/file"

	"../../core"
	"../../storage"
)

// Storage object for file db
type Storage struct {
	name       string
	connection *sql.DB
}

// New create storage instance
func New() *Storage {
	return &Storage{
		name: "file.db",
	}
}

// Init db and create connection. Do migration if needed.
// By default dir = "./local/filestorage/"
func (s *Storage) Init(args ...string) {
	fmt.Println("Init file storage")

	dir := args[0]
	err := os.MkdirAll(dir, os.ModePerm)
	file := dir + "perstorage.db"

	s.connection, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	s.migrate()
}

func (s Storage) migrate() {
	driver, err := sqlite3.WithInstance(s.connection, &sqlite3.Config{})
	if err != nil {
		log.Fatal("can't init sqlite driver", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./storage/filestorage/migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatal("can't get sqlite migration instance", err)
	}
	m.Up()
}

// Close defer db.Close()
func (s Storage) Close() {
	s.connection.Close()
}

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) (int, error) {
	ns, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8") // TODO: namespace
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	key := "salt:" + item.Filename
	u5 := uuid.NewV5(ns, key)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}

	dir := "./local/filestorage/" + u5.String() + "/"
	path := dir + item.Filename

	err = os.MkdirAll(dir, os.ModePerm)
	os.Rename(item.SourceName, path)

	fi, err := os.Stat(path)
	size := int64(-1)
	if err == nil {
		size = fi.Size()
	}

	sql := "INSERT INTO items(name, filename, path, size, available) VALUES(?, ?, ?, ?, ?)"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("", item.Filename, path, size, true)
	if err != nil {
		log.Println("can't insert item into db " + err.Error())
		return 0, err
	}

	if item.IsNew() {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		}
		item.ID = int(id)
	}

	fmt.Println("StoreItem<>")
	return item.ID, err
}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(order core.Order) (int, error) {
	tx, err := s.connection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO orders(customer_id) VALUES(?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(order.Customer.ID)
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

	fmt.Println("StoreOrder<>")
	return order.ID, err
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
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

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(customer core.Customer) (int, error) {
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

	fmt.Println("StoreCustomer<>")
	return customer.ID, err
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
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
		fmt.Println(customer.ID, customer.Phone)
	}
	return customer, ok
}

// StoreCustomerMessenger save chat id to storage by customer
func (s Storage) StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error {
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

func init() {
	storage.Register("file", New())
}
