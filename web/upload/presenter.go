package upload

import "../../messenger/service"

// PlaceOrderResponse response data
type PlaceOrderResponse struct {
	downloadLink string
	phone        string
}

// OnResponse send order message through messenger via registered phone number
func (r PlaceOrderResponse) OnResponse(phone, orderLink string) {
	r.downloadLink = orderLink
	r.phone = phone
	ms := messengers.NewService()
	ms.OrderMessage(phone, orderLink)
}
