package core

// Customer ...
type Customer struct {
	ID    int
	Name  string
	Phone string
}

func (c Customer) IsNew() bool {
	return c.ID == 0
}

// Item ...
type Item struct {
	ID         int
	Name       string
	Value      []byte
	Size       int
	Available  bool
	Filename   string
	SourceName string
}

func (i Item) IsNew() bool {
	return i.ID == 0
}

// Order ...
type Order struct {
	ID       int
	Customer Customer
	Items    []Item
}

func (o Order) IsNew() bool {
	return o.ID == 0
}

// Add file to bucket
func (o *Order) Add(item Item) error {
	o.Items = append(o.Items, item)
	return nil
}

// The MongoDB driver for Go
// https://github.com/globalsign/mgo
