package contracts

// PurgeOrderInput delete order from database if expired
type PurgeOrdersInput interface {
	PurgeOrders()
}
