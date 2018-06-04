package contracts

// PlaceOrderRequest data structure
type PlaceOrderRequest struct {
	Filename   string
	Dir        string
	Phone      string
	Private    bool
	MD5        []byte
	CustomerID int
	//
	Subject     string
	Description string
}

// PlaceOrderInput upload contract
type PlaceOrderInput interface {
	PlaceOrder(request PlaceOrderRequest, output PlaceOrderOutput)
}

// PlaceOrderOutput upload response contract
type PlaceOrderOutput interface {
	OnResponse(phone, orderLink string)
}

// GetSourceName TODO rename filepath?
func (r PlaceOrderRequest) GetSourceName() string {
	return r.Dir + r.Filename
}
