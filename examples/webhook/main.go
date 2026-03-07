package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cloudapp3/tgbot"
	"github.com/cloudapp3/tgbot/ext"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	secret := os.Getenv("WEBHOOK_SECRET")
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	bot, err := tgbot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	app, err := ext.NewApplication(bot)
	if err != nil {
		log.Fatal(err)
	}

	app.AddHandler(ext.NewCommandHandler("start", func(ctx context.Context, c *ext.Context) error {
		msg := c.EffectiveMessage()
		if msg == nil || msg.Chat == nil {
			return nil
		}
		_, err := c.Bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   "welcome",
		})
		return err
	}))

	mux := http.NewServeMux()
	mux.Handle("/telegram/webhook", app.WebhookHandler(secret))

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
