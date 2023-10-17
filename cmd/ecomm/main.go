package main

import (
	app "github.com/capsulated/tutorial.go.interfaces/internal"
)

const NotifierTelegram = "telegram"

type config struct {
	notifier string
}

func main() {
	log := app.NewLogger()

	var notifyService app.Notifier

	cfg := &config{
		notifier: NotifierTelegram,
	}

	if cfg.notifier == NotifierTelegram {
		notifyService = app.NewTelegramNotifyService()
	} else {
		notifyService = app.NewMailNotifyService()
	}

	orderService := app.NewOrderService(log, notifyService)

	orderService.CreateOrder()
}
