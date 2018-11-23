package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Order contains several files which places at ones
type Order struct {
	ID          int
	Customer    Customer
	Items       []Item
	Description string
	Size        int
	Category    Category
	Date        time.Time
}

// IsNew returns true if it is new order
func (o Order) IsNew() bool {
	return o.ID == 0
}

// Add file to bucket
func (o *Order) Add(item Item) error {
	o.Items = append(o.Items, item)
	return nil
}

// Link create order hash based on orderID & customerID
func (o *Order) Link() string {
	key := "order#" + string(o.ID) +
		"customer:" + string(o.Customer.ID)
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func (o *Order) String() string {
	return fmt.Sprintf("Order No:%d %s %s", o.ID, o.Description, o.Date)
}
