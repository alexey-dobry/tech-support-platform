package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/bot/middleware"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) handleEndTicket() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		sender := c.Sender().ID

		if session.IsAuthorized(sender) {
			session.FreeManager(sender)
			log.Printf("Succesfully closed ticket for user %d", sender)
			return c.Send("Тикет успешно закрыт")
		} else {
			log.Printf("Permission denied for command /end user %d", sender)
			return c.Send("Вы не авторизованы для выполнения данной команды")
		}
	}
}

var managerID int64 = 549938415

func (b *bot) HandleGetMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		clientID := c.Sender().ID

		// Проверяем, есть ли у клиента назначенный менеджер
		managerID, err := session.GetAssignedManager(clientID)
		if err != nil {
			return c.Send("Произошла ошибка. Попробуйте позже.")
		}

		if managerID == 0 {
			// Менеджер не назначен, назначаем нового
			managerID, err = session.AssignClientToManager(clientID)
			if err != nil {
				return c.Send("Извините, сейчас нет свободных менеджеров. Пожалуйста, подождите.")
			}
		}

		// Отправляем сообщение назначенному менеджеру
		managerChat := telebot.ChatID(managerID)
		msg := fmt.Sprintf("Новое сообщение от клиента %d: %s", clientID, c.Message().Text)
		_, err = b.Client.Send(managerChat, msg)
		if err != nil {
			log.Println("Ошибка отправки сообщения менеджеру:", err)
			return c.Send("Не удалось доставить сообщение менеджеру. Попробуйте позже.")
		}

		return c.Send("Ваше сообщение отправлено менеджеру.")
	}

}

func (b *bot) HandleAuth() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()
		if len(args) < 2 {
			return c.Send("Используйте команду: /login <логин> <пароль>")
		}

		login := args[0]
		password := args[1]

		// Отправляем запрос к микросервису
		if middleware.Authenticate(login, password) {
			managerID := c.Sender().ID
			if session.IsAuthorized(managerID) {
				return c.Send("Вы уже вошли в аккаунт")
			} else {
				session.AddNewManager(managerID)
				return c.Send("Авторизация успешна! Теперь вы можете отвечать клиентам.")
			}
		}

		return c.Send("Неверный логин или пароль.")
	}
}

func (b *bot) HandleLogut() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()

		if len(args) != 0 {
			return c.Send("Данная команда не пришимает никаких дополнительных значений")
		}

		senderId := c.Sender().ID

		if session.IsAuthorized(senderId) {
			session.DeauthorizeManager(senderId)
			return c.Send("Вы вышли из аккаунта")
		} else {
			return c.Send("Вы не вошли в аккаунт")
		}
	}
}

func (b *bot) HandleSendMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		senderID := c.Sender().ID

		args := c.Args()
		if len(args) == 1 {
			return c.Send("Используйте команду: /reply <сообщение>")
		}

		// Проверяем, является ли отправитель менеджером
		managerClientID, err := session.GetActiveClientForManager(senderID)
		if err != nil {
			return c.Send("У вас нет активной сессии.")
		}

		if managerClientID == 0 {
			return c.Send("Нет активных клиентов для ответа.")
		}

		// Пересылаем сообщение клиенту
		clientChat := telebot.ChatID(managerClientID)
		msg := strings.Join(args[:], " ")
		_, err = b.Client.Send(clientChat, msg)
		if err != nil {
			log.Println("Ошибка отправки клиенту:", err)
			return c.Send("Не удалось доставить сообщение клиенту.")
		}

		return c.Send("Ваше сообщение отправлено клиенту.")
	}
}
