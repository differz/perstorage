package contracts

// PurgeOrdersInput delete order from database if expired
type PurgeOrdersInput interface {
	PurgeOrders()
}
