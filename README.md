# tgbot

[![CI](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22%2B-00ADD8?logo=go)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/cloudapp3/tgbot.svg)](https://pkg.go.dev/github.com/cloudapp3/tgbot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/cloudapp3/tgbot/blob/master/LICENSE)

[English](./README.md) | [简体中文](./docs/readme/README.zh-CN.md) | [Español](./docs/readme/README.es.md) | [日本語](./docs/readme/README.ja.md)

A lightweight, strongly typed Telegram Bot API SDK for Go with zero third-party runtime dependencies.

`tgbot` is an open-source Go SDK for building Telegram bots with clean abstractions, full Bot API coverage, typed union decoding, webhook support, and production-friendly long polling.

## Why tgbot?

- Full Telegram Bot API method and type coverage
- Strongly typed union decoding for polymorphic Telegram fields and results
- Zero third-party runtime dependencies (standard library only)
- File upload helpers for `file_id`, URL, local path, and `io.Reader`
- Async long polling with channel-based subscriptions, `Unsubscribe()`, and optional non-blocking dispatch
- PTB-style routing layer in `ext` for:
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

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

## Install

```bash
go get github.com/cloudapp3/tgbot
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
