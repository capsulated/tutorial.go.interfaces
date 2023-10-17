package internal

type OrderService struct {
	log           *Logger
	notifyService CompositedNotifier
}

func NewOrderService(log *Logger, ns CompositedNotifier) *OrderService {
	return &OrderService{
		log:           log,
		notifyService: ns,
	}
}

func (os *OrderService) CreateOrder() {}
