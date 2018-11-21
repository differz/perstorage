package contracts

// PurgeOrderInput delete order from database if expired
type PurgeOrderInput interface {
	PurgeOrder()
}
