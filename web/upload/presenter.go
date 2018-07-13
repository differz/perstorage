package upload

import "../../messenger/service"

// PlaceOrderResponse response data
type PlaceOrderResponse struct {
	ms messengers.Service
}

// OnResponse send order message through messenger via registered phone number
func (r PlaceOrderResponse) OnResponse(phone, orderLink, description string) {
	r.ms.OrderMessage(phone, orderLink, description)
}
