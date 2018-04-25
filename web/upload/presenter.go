package upload

import "../../messenger/service"

// PlaceOrderResponse ...
type PlaceOrderResponse struct {
	downloadLink string
	phone        string
}

// OnResponse ...
func (r PlaceOrderResponse) OnResponse(phone, orderLink string) {
	r.downloadLink = orderLink
	r.phone = phone
	ms := messengers.NewService()
	ms.OrderMessage(phone, orderLink)
}
