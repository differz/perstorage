package core

// Customer ...
type Customer struct {
	ID    int
	Name  string
	Phone string
}

// Item ...
type Item struct {
	ID        int
	Name      string
	Value     []byte
	Size      int
	Available bool
	Filename  string
}

// Order ...
type Order struct {
	ID       int
	Customer Customer
	Items    []Item
}

// Add file to bucket
func (o *Order) Add(item Item) error {
	o.Items = append(o.Items, item)
	return nil
}

// The MongoDB driver for Go
// https://github.com/globalsign/mgo
