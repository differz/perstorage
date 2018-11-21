package contracts

// OrderMessageInput delivery order link to messenger contract
type OrderMessageInput interface {
	OrderMessage(phone, orderLink, description string)
}
