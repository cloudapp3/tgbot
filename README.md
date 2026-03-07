# tgbot

[English](./README.md) | [简体中文](./docs/readme/README.zh-CN.md) | [Español](./docs/readme/README.es.md) | [日本語](./docs/readme/README.ja.md)

A lightweight Go Telegram Bot API SDK with full update routing, webhook handler support, and clean, production-friendly abstractions.

## Features

- Full Telegram Bot API method/type coverage included in this repository, verified against the official docs
- Typed union decoding for Telegram polymorphic fields/results
- Zero third-party runtime dependencies (standard library only)
- File upload helpers (`file_id`, URL, local path, `io.Reader`)
- Async long-polling helper with channel-based subscriptions, `Unsubscribe()`, and optional non-blocking dispatch on top of `GetUpdates`
- PTB-style routing layer in `ext`:
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## Install

```bash
go get github.com/cloudapp3/tgbot
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/cloudapp3/tgbot"
)

func main() {
    bot, err := tgbot.NewBot("<BOT_TOKEN>")
    if err != nil {
        log.Fatal(err)
    }

    _, err = bot.SendMessage(context.Background(), &tgbot.SendMessageParams{
        ChatID: int64(123456789),
        Text:   "hello from tgbot",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

## Webhook Routing (`ext`)

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
            Text:   "welcome",
        })
        return err
    }))

    mux := http.NewServeMux()
    mux.Handle("/telegram/webhook", app.WebhookHandler("<SECRET_TOKEN>"))
    _ = http.ListenAndServe(":8080", mux)
}
```

## Examples

- `examples/quickstart`
- `examples/webhook`
- `examples/commands`
- `examples/async_updates`
- `examples/ext_polling`

## Official Coverage

This SDK is verified against the official Telegram Bot API documentation at `https://core.telegram.org/bots/api`.
As of `2026-03-06`, the latest official release on that page is `Bot API 9.5` from `March 1, 2026`.

```bash
python3 scripts/verify_telegram_bot_api_coverage.py
```

## Sync SDK Definitions

```bash
go run ./cmd/apigen
```

If you already downloaded the official HTML page:

```bash
go run ./cmd/apigen -html /tmp/telegram_bot_api.html
```

Key files:

- `sdk_types.go`
- `sdk_methods.go`
- `sdk_unions.go`

## Development

```bash
go test ./...
go vet ./...
```

## License

MIT
