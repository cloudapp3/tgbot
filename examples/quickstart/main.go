package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudapp3/tgbot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")
	if token == "" || chatID == "" {
		log.Fatal("BOT_TOKEN and CHAT_ID are required")
	}

	bot, err := tgbot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.SendMessage(context.Background(), &tgbot.SendMessageParams{
		ChatID: chatID,
		Text:   "hello from tgbot quickstart",
	})
	if err != nil {
		log.Fatal(err)
	}
}
