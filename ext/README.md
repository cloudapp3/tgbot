# ext package quick start

`ext` provides PTB-like app-layer routing on top of `tgbot.Bot`:

- `Application`
- `Handler`
- `Filter`
- webhook `http.Handler`
- long-polling runner via `Application.RunPolling`
- optional non-blocking polling dispatch via `WithPollingNonBlockingDispatch()`

## Minimal Setup

```go
package main

import (
    "context"
    "net/http"

    "github.com/cloudapp3/tgbot"
    "github.com/cloudapp3/tgbot/ext"
)

func main() {
    bot, _ := tgbot.NewBot("<BOT_TOKEN>")
    app, _ := ext.NewApplication(bot)

    app.AddHandler(ext.NewCommandHandler("start", func(ctx context.Context, c *ext.Context) error {
        msg := c.EffectiveMessage()
        if msg == nil || msg.Chat == nil {
            return nil
        }
        _, err := c.Bot.SendMessage(ctx, &tgbot.SendMessageParams{
            ChatID: msg.Chat.ID,
            Text:   "hello",
        })
        return err
    }))

    http.Handle("/telegram/webhook", app.WebhookHandler("<SECRET_TOKEN>"))
    _ = http.ListenAndServe(":8080", nil)
}
```

## Long Polling

```go
package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"

    "github.com/cloudapp3/tgbot"
    "github.com/cloudapp3/tgbot/ext"
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    bot, _ := tgbot.NewBot("<BOT_TOKEN>")
    app, _ := ext.NewApplication(bot)

    app.AddHandler(ext.NewCommandHandler("start", func(ctx context.Context, c *ext.Context) error {
        return nil
    }))

    _ = app.RunPolling(ctx,
        ext.WithPollingAllowedUpdates(
            ext.UpdateTypeMessage,
            ext.UpdateTypeCallbackQuery,
        ),
        ext.WithPollingNonBlockingDispatch(),
    )
}
```

## Useful Helpers

- Handlers:
  - `NewAnyHandler`
  - `NewTypeHandler`
  - `NewMessageHandler`
  - `NewCommandHandler`
  - `NewCallbackQueryHandler`
- Filters:
  - `TextFilter`
  - `CommandFilter`
  - `RegexFilter`
  - `UpdateTypeFilter`
  - `And/Or/Not`
