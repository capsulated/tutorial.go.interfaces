package internal

type MailNotifyService struct {
	Config       *Config
	CreatedTpl   *Template
	PaidTpl      *Template
	CompletedTpl *Template
}

func NewMailNotifyService() *MailNotifyService {
	return &MailNotifyService{}
}

func (ns *MailNotifyService) OrderCreated(order *Order) error {
	// ..
	return nil
}

func (ns *MailNotifyService) OrderPaid(order *Order) (int, error) {
	// ..
	return 0, nil
}

func (ns *MailNotifyService) OrderCompleted(order *Order) error {
	// ..
	return nil
}
