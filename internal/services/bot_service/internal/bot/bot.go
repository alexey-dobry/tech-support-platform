package bot

import (
	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger"
	"gopkg.in/telebot.v4"
)

type Bot interface {
	Run()
}

type bot struct {
	middlewareAddress string
	Client            *telebot.Bot
	Logger            logger.Logger
}

// Создание инстанса бота
func New(client *telebot.Bot, logger logger.Logger, port string) Bot {
	var b bot

	b.Logger = logger

	b.Client = client

	b.middlewareAddress = port

	b.initHandlers()

	return &b
}

// Инициализация функций бота
func (b *bot) initHandlers() {
	// Клиенты отправляют сообщения
	b.Client.Handle(telebot.OnText, b.HandleGetMsg())

	// Менеджер отвечает клиенту
	b.Client.Handle("/reply", b.HandleSendMsg())

	// Вход в систему(для менеджера)
	b.Client.Handle("/login", b.HandleAuth())

	// Выход из системы(для менеджера)
	b.Client.Handle("/logout", b.HandleLogut())

	b.Client.Handle("/end", b.handleEndTicket())
}

// Запуск бота
func (b *bot) Run() {
	b.Client.Start()
}
