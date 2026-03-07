package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudapp3/tgbot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	bot, err := tgbot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	var offset int64

	for {
		updates, err := bot.GetUpdates(ctx, &tgbot.GetUpdatesParams{
			Offset:  offset,
			Timeout: 25,
			Limit:   100,
		})
		if err != nil {
			log.Printf("getUpdates failed: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, update := range updates {
			if update.UpdateID >= offset {
				offset = update.UpdateID + 1
			}
			if update.Message == nil || update.Message.Chat == nil {
				continue
			}

			command, args, ok := tgbot.ParseCommand(update.Message.Text)
			if !ok {
				continue
			}

			reply := "unknown command"
			switch command {
			case "ping":
				reply = "pong"
			case "echo":
				reply = args
			}

			_, err := bot.SendMessage(ctx, &tgbot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   reply,
			})
			if err != nil {
				log.Printf("sendMessage failed: %v", err)
			}
		}
	}
}
