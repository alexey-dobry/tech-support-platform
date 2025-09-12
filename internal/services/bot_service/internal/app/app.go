package app

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/bot"
)

type App interface {
	Run()
}

type app struct {
	Bot bot.Bot
}

func New(bot bot.Bot) App {
	var a app

	a.Bot = bot

	log.Print("The app is built")
	return &a
}

func (a *app) Run() {
	log.Print("The app is running...")
	a.Bot.Run()
}
