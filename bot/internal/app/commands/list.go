package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	list := ""
	for _, v := range c.productService.ListOfAllProducts() {
		list += fmt.Sprintf("%s\n", v.Title)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, list)
	c.bot.Send(msg)
}
