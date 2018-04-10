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
		name: "file.txt",
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
func (s Storage) StoreItem(item core.Item) {
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

	item.ID = int(crc32.ChecksumIEEE([]byte(key)))

	fmt.Println(err)

	fileDB := dir + "perstorage.db"

	//	os.Remove(fileDB)

	db, err := sql.Open("sqlite3", fileDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./storage/filestorage/migrations",
		"sqlite3", driver)
	m.Up()

}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(item core.Order) {
	fmt.Println("StoreOrder<>")
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
	return core.Order{}, false
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) {
	fmt.Println("StoreCustomer<>")
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
	return core.Customer{}, false
}

func init() {
	storage.Register("file", New())
}
