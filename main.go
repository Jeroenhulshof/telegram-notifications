package main

import (
	"log"
)

func main() {

	bot := TelegramBot{
		Token: getToken(),
	}

	log.Print(bot.getMe())
}
