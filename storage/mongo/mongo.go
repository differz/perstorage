package mongo

// The MongoDB driver for Go
// https://github.com/globalsign/mgo
import (
	"github.com/differz/perstorage/common"
	"github.com/differz/perstorage/core"
	"github.com/differz/perstorage/storage"
)

// Storage object for mongo db
type Storage struct {
	name string
}

const component = "mongo"

// New create storage instance
func New() Storage {
	return Storage{
		name: "MongoDB",
	}
}

// Init db and create connection. Do migration if needed.
func (s Storage) Init(args ...string) {
	common.ContextUpMessage(component, "init mongo storage")
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

// DeleteItem remove file from storage
func (s Storage) DeleteItem(item core.Item) bool {
	return false
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

// GetOrders takes oreders from db
func (s Storage) GetOrders(strategy func() string) ([]core.Order, error) {
	return nil, nil
}

// GetOrderedItems takes all ordered items by selected order
func (s Storage) GetOrderedItems(order core.Order) ([]core.Item, bool) {
	return nil, false
}

// DeleteOrder remove order from storage
func (s Storage) DeleteOrder(order core.Order) bool {
	return false
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

// IsRegisteredChatID get customer's id if chat already registered
func (s Storage) IsRegisteredChatID(chatID int, messengerName string) (int, bool) {
	return 0, false
}

func (s Storage) String() string {
	return s.name
}

func init() {
	storage.Register(component, New())
}
