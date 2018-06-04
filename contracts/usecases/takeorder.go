package contracts

// TakeOrderRequest data structure
type TakeOrderRequest struct {
	Link string
}

// TakeOrderInput download contract
type TakeOrderInput interface {
	TakeOrder(request TakeOrderRequest, output TakeOrderOutput)
}

// TakeOrderOutput download response contract
type TakeOrderOutput interface {
	OnResponse(sourcename, filename string, size int)
}
