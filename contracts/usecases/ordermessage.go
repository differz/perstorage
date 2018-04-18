package contracts

type OrderMessageRequest struct {
	Phone string
}

type OrderMessageInput interface {
	OrderMessage(request OrderMessageRequest, output OrderMessageOutput)
}

type OrderMessageOutput interface {
	OnResponse(orderLink string)
}
