package contracts

type TakeOrderRequest struct {
	Filename string
	Phone    string
	Private  bool
	//
	OrderId     int
	Subject     string
	Description string
}

type TakeOrderInput interface {
	TakeOrder(request TakeOrderRequest, output TakeOrderOutput)
}

type TakeOrderOutput interface {
	OnResponse(orderID int)
}
