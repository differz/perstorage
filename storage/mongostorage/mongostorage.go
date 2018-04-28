package mongostorage

// The MongoDB driver for Go
// https://github.com/globalsign/mgo
import (
	"fmt"

	"../../core"
	"../../storage"
)

// Storage object for mongo db
type Storage struct {
	name string
}

// New create storage instance
func New() Storage {
	return Storage{
		name: "MongoDB",
	}
}

// Init db and create connection. Do migration if needed.
func (s Storage) Init(args ...string) {
	fmt.Println("Init<>")
	s.migrate()
}

func (s Storage) migrate() {

}

// Close defer db.Close()
func (s Storage) Close() {

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

// StoreCustomerMessenger save chat id to storage by customer
func (s Storage) StoreCustomerMessenger(customer core.Customer, messengerName string, chatID int) error {
	return nil
}

// FindCustomerChatID  get customer's chat id
func (s Storage) FindCustomerChatID(customer core.Customer, messengerName string) (int, bool) {
	return 0, false
}

func init() {
	storage.Register("mongo", New())
}
