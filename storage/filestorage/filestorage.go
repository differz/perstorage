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

// Storage ...
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

// InitDB ...
func (s *Storage) InitDB() *sql.DB {
	fmt.Println("InitDB file storage")

	dir := "./local/filestorage/"
	err := os.MkdirAll(dir, os.ModePerm)
	fileDB := dir + "perstorage.db"

	s.connection, err = sql.Open("sqlite3", fileDB)
	if err != nil {
		log.Fatal(err)
	}

	// db, _ := gorm.Open("sqlite3", "test.db")
	// db.Exec("PRAGMA foreign_keys = ON")
	// db.LogMode(true)
	// db.AutoMigrate(&User{}, &Address{})
	// fmt.Println(db.Save(&User{}).Error)
	// fmt.Println(db.Save(&Address{}).Error)

	//	defer db.Close()
	return s.connection
}

// Migrate ...
func (s Storage) Migrate() {
	driver, err := sqlite3.WithInstance(s.connection, &sqlite3.Config{})
	//TODO: if err != nil {}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./storage/filestorage/migrations",
		"sqlite3", driver)
	if err != nil {
		return
	}
	m.Up()
}

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) (int, error) {

	key := "salt:" + item.Filename

	// TODO: namespace
	ns, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
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
		panic("Exec err:" + err.Error())
		//TODO error
	}

	if item.IsNew() {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		} else {
			println("LastInsertId:", id)
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
		} else {
			println("LastInsertId:", id)
		}
		order.ID = int(id)
	}

	sql = "INSERT INTO ordered_items(order_id, customer_id, item_id, link) VALUES(?, ?, ?, ?)"
	stmt, err = tx.Prepare(sql)
	for index, item := range order.Items {
		fmt.Println(index)
		_, err := stmt.Exec(order.ID, order.Customer.ID, item.ID, order.Link())
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
	return core.Order{}, false
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
		} else {
			println("LastInsertId:", id)
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

func init() {
	storage.Register("file", New())
}
