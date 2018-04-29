package core

import (
	"strconv"
	"strings"
)

// Customer object identifies customer by phone number
type Customer struct {
	ID    int
	Name  string
	Phone string
}

// IsNew returns true if it is new customer
func (c Customer) IsNew() bool {
	return c.ID == 0
}

// GetCustomerIDByPhone takes customer id by phone number
func GetCustomerIDByPhone(phone string) (int, error) {
	return strconv.Atoi(strings.Replace(phone, "+", "", 1))
}
