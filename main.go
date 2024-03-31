package main

import (
	"flag"
)

func main() {
	token := flag.String("token", "Something wen't wrong", "Telegram Token")
	chatId := flag.String("chatId", "0", "Chat ID")
	message := flag.String("message", "Something wen't wrong", "Message")
	flag.Parse()

	bot := TelegramBot{
		Token: *token,
	}

	bot.sendMessage(*chatId, *message)
}
