package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudapp3/tgbot"
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

	poller, err := bot.StartUpdatePoller(ctx,
		tgbot.WithPollerParams(tgbot.GetUpdatesParams{
			Timeout: 25,
			Limit:   100,
		}),
		tgbot.WithPollerNonBlockingDispatch(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer poller.Stop()

	messageSub := poller.SubscribeTypesHandle(32, tgbot.UpdateTypeMessage, tgbot.UpdateTypeEditedMessage)
	defer messageSub.Unsubscribe()
	messageCh := messageSub.Updates()

	callbackSub := poller.SubscribeTypesHandle(16, tgbot.UpdateTypeCallbackQuery)
	defer callbackSub.Unsubscribe()
	callbackCh := callbackSub.Updates()

	go func() {
		for err := range poller.Errors() {
			log.Printf("poller error: %v", err)
		}
	}()

	go func() {
		for update := range callbackCh {
			if update.CallbackQuery == nil {
				continue
			}
			log.Printf("callback query: %s", update.CallbackQuery.Data)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case update, ok := <-messageCh:
			if !ok {
				return
			}
			message := update.EffectiveMessage()
			if message == nil || message.Chat == nil {
				continue
			}

			command, args, ok := update.Command()
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
				ChatID: message.Chat.ID,
				Text:   reply,
			})
			if err != nil {
				log.Printf("sendMessage failed: %v", err)
			}
		}
	}
}
