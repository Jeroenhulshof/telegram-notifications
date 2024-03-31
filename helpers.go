package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func newGetRequest(apiEndpoint string) (string, error) {
	response, error := http.Get(apiEndpoint)

	if error != nil {
		log.Fatal("Error making GET request:", error)
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal("Error reading response body:", error)
	}

	return string(body), nil
}

func newPostRequest(apiEndpoint string, chatId string, message string) {
	payload := strings.NewReader(fmt.Sprintf(`{"chat_id":"%s", "text":"%s", "parse_mode":"Markdown", "disable_web_page_preview":false, "disable_notification":false, "reply_to_message_id":null}`, chatId, message))

	request, _ := http.NewRequest("POST", apiEndpoint, payload)

	request.Header.Add("accept", "application/json")
	request.Header.Add("User-Agent", "Telegram Bot SDK - (https://github.com/irazasyed/telegram-bot-sdk)")
	request.Header.Add("content-type", "application/json")

	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()

	io.ReadAll(response.Body)
}
