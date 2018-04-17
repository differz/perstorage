package core

import (
	"crypto/sha256"
	"encoding/hex"
)

// Order ...
type Order struct {
	ID       int
	Customer Customer
	Items    []Item
}

// IsNew ...
func (o Order) IsNew() bool {
	return o.ID == 0
}

// Add file to bucket
func (o *Order) Add(item Item) error {
	o.Items = append(o.Items, item)
	return nil
}

// Link ...
func (o *Order) Link() string {
	key := "order#" + string(o.ID) + ", customer:" + string(o.Customer.ID)
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)
	return hashHex
}
