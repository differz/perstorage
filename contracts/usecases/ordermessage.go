package contracts

// OrderMessageInput delivery order link to messenger contract
type OrderMessageInput interface {
	OrderMessage(phone string, orderLink, description string)
}

/*
type OrderMessageRequest struct {
	Phone   string
	Message string
}

type OrderMessageInput interface {
	OrderMessage(request OrderMessageRequest, output OrderMessageOutput)
}

type OrderMessageOutput interface {
	OnResponse()
}
*/
