package contracts

type TakeOrderRequest struct {
	Link string
}

type TakeOrderInput interface {
	TakeOrder(request TakeOrderRequest, output TakeOrderOutput)
}

type TakeOrderOutput interface {
	OnResponse(orderID int)
}
