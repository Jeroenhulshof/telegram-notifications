package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TelegramBot struct {
	Token string
}

func (bot TelegramBot) getMe() (string, error) {
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.Token)

	return newGetRequest(apiEndpoint)
}

func (bot TelegramBot) sendMessage(chatId string, message string) {
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.Token)

	payload := strings.NewReader(fmt.Sprintf(`{"chat_id":"%s", "text":"%s", "parse_mode":"Markdown", "disable_web_page_preview":false, "disable_notification":false, "reply_to_message_id":null}`, chatId, message))

	request, _ := http.NewRequest("POST", apiEndpoint, payload)

	request.Header.Add("accept", "application/json")
	request.Header.Add("User-Agent", "Telegram Bot SDK - (https://github.com/irazasyed/telegram-bot-sdk)")
	request.Header.Add("content-type", "application/json")

	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()

	io.ReadAll(response.Body)
}
