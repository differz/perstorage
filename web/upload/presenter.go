package upload

import "../../messenger/service"

// PlaceOrderResponse response data
type PlaceOrderResponse struct {
	downloadLink string
	phone        string
	ms           messengers.Service
}

// OnResponse send order message through messenger via registered phone number
func (r PlaceOrderResponse) OnResponse(phone, orderLink string) {
	r.downloadLink = orderLink
	r.phone = phone
	r.ms.OrderMessage(phone, orderLink)
}
