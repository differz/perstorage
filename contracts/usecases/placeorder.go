package contracts

type PlaceOrderRequest struct {
	Filename string
	//
	OrderId     int64
	Subject     string
	Description string
}

type PlaceOrderInput interface {
	PlaceOrder(request PlaceOrderRequest, output PlaceOrderOutput)
}

type PlaceOrderOutput interface {
	OnResponse(orderID int64)
}
