package contracts

/*
type OrderMessageRequest struct {
	Phone   string
	Message string
}

type OrderMessageInput interface {
	OrderMessage(request OrderMessageRequest, output OrderMessageOutput)
}

type OrderMessageOutput interface {
	OnResponse(phone string, orderLink string)
}
*/

type OrderMessageInput interface {
	OrderMessage(phone string, orderLink string)
}
