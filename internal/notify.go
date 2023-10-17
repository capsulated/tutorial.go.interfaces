package internal

type Config struct{}

type Template struct{}

type NotifyService struct {
	Config       *Config
	CreatedTpl   *Template
	PaidTpl      *Template
	CompletedTpl *Template
}

func (ns *NotifyService) OrderCreated(order *Order) error {
	// ..
	return nil
}

func (ns *NotifyService) OrderPaid(order *Order) (int, error) {
	// ..
	return 0, nil
}

func (ns *NotifyService) OrderCompleted(order *Order) error {
	// ..
	return nil
}
