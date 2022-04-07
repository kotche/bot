package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kotche/bot/internal/app/commands"
	"github.com/kotche/bot/internal/service/product"
)

func main() {

	//Do not process the error, because want to work 'TOKEN=123.. make run' , if not set TOKEN=123.. in ".env"
	godotenv.Load()

	//Get token from environment variable
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//Log output to terminal
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
