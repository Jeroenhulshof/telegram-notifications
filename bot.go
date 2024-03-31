package main

import (
	"fmt"
	"log"
)

type TelegramBot struct {
	Token string
}

func (bot TelegramBot) getMe() (string, error) {
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.Token)

	return newGetRequest(apiEndpoint)
}

func (bot TelegramBot) sendMessage(chatId string, message string) {
	log.Print(bot.Token)
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.Token)

	newPostRequest(apiEndpoint, chatId, message)
}
