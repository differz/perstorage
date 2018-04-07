package contracts

type PlaceOrderRequest struct {
	Filename string
	Phone    string
	Private  bool
	//
	OrderId     int
	Subject     string
	Description string
}

type PlaceOrderInput interface {
	PlaceOrder(request PlaceOrderRequest, output PlaceOrderOutput)
}

type PlaceOrderOutput interface {
	OnResponse(orderID int)
}
