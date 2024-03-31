package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	bot := TelegramBot{
		Token: getToken(),
	}

	log.Print(bot.getMe())
}

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
