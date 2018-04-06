package filestorage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

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

	fmt.Println(err)
}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) core.Item {
	return core.Item{}
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(item core.Order) {
	fmt.Println("StoreOrder<>")
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) core.Order {
	return core.Order{}
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) {
	fmt.Println("StoreCustomer<>")
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) core.Customer {
	return core.Customer{}
}

func init() {
	storage.Register("file", New())
}
