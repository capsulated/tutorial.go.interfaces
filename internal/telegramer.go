package internal

type TgBotApi struct{}

type TelegramNotifyService struct {
	TgBotApi     *TgBotApi
	CreatedMsg   string
	PaidMsg      string
	CompletedMsg string
}

func NewTelegramNotifyService() *TelegramNotifyService {
	return &TelegramNotifyService{}
}

func (ns *TelegramNotifyService) OrderCreated(order *Order) error {
	// ..
	return nil
}

func (ns *TelegramNotifyService) OrderPaid(order *Order) (int, error) {
	// ..
	return 0, nil
}

func (ns *TelegramNotifyService) OrderCompleted(order *Order) error {
	// ..
	return nil
}
