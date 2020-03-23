package contracts

// PlaceOrderRequest data structure
type PlaceOrderRequest struct {
	Filename    string
	Dir         string
	Phone       string
	Private     bool
	MD5         []byte
	CustomerID  int
	CategoryID  int
	Description string
}

// PlaceOrderInput upload contract
type PlaceOrderInput interface {
	PlaceOrder(request PlaceOrderRequest, output PlaceOrderOutput)
}

// PlaceOrderOutput upload response contract
type PlaceOrderOutput interface {
	OnResponse(phone, orderLink, description string)
}

// GetFullFileName Dir/Filename
func (r PlaceOrderRequest) GetFullFileName() string {
	return r.Dir + r.Filename
}
