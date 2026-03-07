package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudapp3/tgbot"
	"github.com/cloudapp3/tgbot/ext"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	bot, err := tgbot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	app, err := ext.NewApplication(bot)
	if err != nil {
		log.Fatal(err)
	}

	app.AddHandler(ext.NewCommandHandler("start", func(ctx context.Context, c *ext.Context) error {
		message := c.EffectiveMessage()
		if message == nil || message.Chat == nil {
			return nil
		}
		_, err := c.Bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: message.Chat.ID,
			Text:   "welcome from polling",
		})
		return err
	}))

	app.AddHandler(ext.NewCallbackQueryHandler(nil, func(ctx context.Context, c *ext.Context) error {
		if c == nil || c.Update == nil || c.Update.CallbackQuery == nil {
			return nil
		}
		log.Printf("callback query: %s", c.Update.CallbackQuery.Data)
		return nil
	}))

	log.Printf("polling started")
	if err := app.RunPolling(ctx,
		ext.WithPollingAllowedUpdates(
			ext.UpdateTypeMessage,
			ext.UpdateTypeEditedMessage,
			ext.UpdateTypeCallbackQuery,
		),
		ext.WithPollingNonBlockingDispatch(),
	); err != nil {
		log.Fatal(err)
	}
}
