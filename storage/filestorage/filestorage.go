package filestorage

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"log"
	"os"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"

	"../../configuration"
	"../../core"
	"../../storage"
)

// Storage ...
type Storage struct {
	name string
}

// New create storage instance
func New() Storage {
	return Storage{
		name: "file.db",
	}
}

// InitDB ...
func (s Storage) InitDB() *sql.DB {
	fmt.Println("InitDB<>")

	dir := "./local/filestorage/"
	err := os.MkdirAll(dir, os.ModePerm)
	fileDB := dir + "perstorage.db"

	db, err := sql.Open("sqlite3", fileDB)
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
	return db
}

// Migrate ...
func (s Storage) Migrate(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
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
	fmt.Println("StoreItem<>")

	key := "salt:" + item.Filename
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)

	// create <dir> 3aa16ca782e477f5ae798e68c5b335b4e77873bbf9154ac5fd0fa098ef2b1c51
	dir := "./local/filestorage/"

	path := dir + hashHex + "/" + item.Filename

	err := os.MkdirAll(dir+hashHex, os.ModePerm)
	os.Rename("./local/"+item.Filename, path)

	fi, err := os.Stat(path)
	size := int64(-1)
	if err == nil {
		size = fi.Size()
	}

	fmt.Println(err)

	sql := "INSERT INTO items(name, filename, path, size, available) VALUES(?, ?, ?, ?, ?)"
	db := configuration.Get().Connection
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec("", item.Filename, path, size, true)

	itemID := int(crc32.ChecksumIEEE([]byte(key)))
	return itemID, nil
}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(order core.Order) (int, error) {
	sql := "INSERT INTO orders(order_id, customer_id, item_id) VALUES(?, ?, ?)"
	db := configuration.Get().Connection
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for index, item := range order.Items {
		fmt.Println(index)
		_, err = stmt.Exec(order.ID, order.Customer.ID, item.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	fmt.Println("StoreOrder<>")
	return 0, nil
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
	return core.Order{}, false
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) (int, error) {
	sql := "INSERT INTO customers(id, name, phone) VALUES(?, ?, ?)"
	db := configuration.Get().Connection
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(item.ID, item.Name, item.Phone)

	fmt.Println("StoreCustomer<>")
	return 0, nil
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
	sql := "SELECT id, name, phone FROM customers WHERE id = ?"
	db := configuration.Get().Connection
	stmt, err := db.Prepare(sql)
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
