package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Commander) handleMessage(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		//c.Get(msg)
	default:
		c.Default(msg)
	}
}
