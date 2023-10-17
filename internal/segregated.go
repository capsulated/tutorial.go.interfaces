package internal

type OrderCreatedNotifier interface {
	OrderCreated(order *Order) error
}

type OrderPaidNotifier interface {
	OrderCreated(order *Order) error
}

type OrderCompletedNotifier interface {
	OrderCreated(order *Order) error
}

type CompositedNotifier interface {
	OrderCreatedNotifier
	OrderPaidNotifier
	OrderCompletedNotifier
}
