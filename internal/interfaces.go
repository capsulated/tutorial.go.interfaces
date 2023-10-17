package internal

type Notifier interface {
	OrderCreated(order *Order) error
	OrderPaid(order *Order) (int, error)
	OrderCompleted(order *Order) error
}
