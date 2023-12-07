package main

import (
	"/enrollments-parser/config"
	"/pkg/bot"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func main() {
	config.Init()

	botApi, err := tgbotapi.NewBotAPI(viper.GetString("tgbot_token"))
	if err != nil {
		log.Fatal(err)
	}

	bfBot := bot.NewBot(botApi)
	err = bfBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
