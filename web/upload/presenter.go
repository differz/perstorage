package upload

import "../../messenger"

// PlaceOrderResponse ...
type PlaceOrderResponse struct {
	downloadLink string
}

// OnResponse ...
func (r PlaceOrderResponse) OnResponse(orderLink string) {
	r.downloadLink = orderLink
	ms := messenger.NewService()
	_ = ms
}
