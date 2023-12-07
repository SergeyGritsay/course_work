package commands

import (
	"/enrollments-parser/pkg/bot/response"
	"/enrollments-parser/pkg/bot/strategies"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

type CommandStrategy interface {
	strategies.HandleStrategy
	getCommandName() string
}

type StartStrategy struct{}

func NewStartStrategy() *StartStrategy {
	return &StartStrategy{}
}

func (s *StartStrategy) Handle(update *tgbotapi.Update) (response.BotResponse, error) {
	return response.NewInstantResponse(
		update.Message.Chat.ID,
		viper.GetString("commands.start"),
		response.TextMessage), nil
}
