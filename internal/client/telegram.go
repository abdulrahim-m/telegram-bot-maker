package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramClient(token string) (*TelegramClient, error) {
	// Initialize the bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramClient{
		bot: bot,
	}, nil
}
