package main

import (
	"flag"
	"fmt"
)

type TelegramBot struct {
	Token string
}

func getToken() string {
	token := flag.String("token", "value", "Fetching token")

	flag.Parse()

	return *token
}

func (bot TelegramBot) getMe() (string, error) {
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.Token)

	return newGetRequest(apiEndpoint)
}
