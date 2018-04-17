package mongostorage

// The MongoDB driver for Go
// https://github.com/globalsign/mgo
import (
	"database/sql"
	"fmt"

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
		name: "MongoDB",
	}
}

// InitDB ...
func (s Storage) InitDB() *sql.DB {
	fmt.Println("InitDB<>")
	return nil
}

// Migrate ...
func (s Storage) Migrate() {

}

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) (int, error) {
	return 0, nil
}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// StoreOrder save bucket to storage
func (s Storage) StoreOrder(item core.Order) (int, error) {
	return 0, nil
}

// FindOrderByID get bucket from storage
func (s Storage) FindOrderByID(id int) (core.Order, bool) {
	return core.Order{}, false
}

// FindOrderByLink get bucket from storage by link
func (s Storage) FindOrderByLink(link string) (core.Order, bool) {
	return core.Order{}, false
}

// StoreCustomer save client to storage
func (s Storage) StoreCustomer(item core.Customer) (int, error) {
	return 0, nil
}

// FindCustomerByID get client from storage
func (s Storage) FindCustomerByID(id int) (core.Customer, bool) {
	return core.Customer{}, false
}

func init() {
	storage.Register("mongo", New())
}
