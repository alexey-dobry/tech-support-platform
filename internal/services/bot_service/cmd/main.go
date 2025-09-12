package main

import (
	"time"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger/zap"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/app"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/bot"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/config"
	"gopkg.in/telebot.v4"
)

func main() {
	cfg := config.MustLoad()

	logger := zap.NewLogger(cfg.Logger)

	client, _ := telebot.NewBot(telebot.Settings{
		Token:  cfg.Bot.Tocken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	bot := bot.New(client, logger)

	app := app.New(bot)

	app.Run()
}
