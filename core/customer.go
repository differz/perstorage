package core

// Customer ...
type Customer struct {
	ID    int
	Name  string
	Phone string
}

// IsNew ...
func (c Customer) IsNew() bool {
	return c.ID == 0
}
