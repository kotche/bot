package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kotche/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	//Call before closing HandleUpdate. Prevents the program from crashing
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	//Example parse string and json
	if update.CallbackQuery != nil {

		var text string

		if strings.Contains(update.CallbackQuery.Data, "_") {
			args := strings.Split(update.CallbackQuery.Data, "_")

			text = fmt.Sprintf("Command: %s\n", args[0]) +
				fmt.Sprintf("Offset: %s\n", args[1])

		} else {
			parsedData := CommandData{}
			json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
			text = fmt.Sprintf("Parsed data: %+v\n", parsedData)
		}

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)

		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}
	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
