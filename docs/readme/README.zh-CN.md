# tgbot

[![CI](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22%2B-00ADD8?logo=go)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/cloudapp3/tgbot.svg)](https://pkg.go.dev/github.com/cloudapp3/tgbot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/cloudapp3/tgbot/blob/master/LICENSE)

[English](../../README.md) | [简体中文](./README.zh-CN.md) | [Español](./README.es.md) | [日本語](./README.ja.md)

轻量、强类型、零运行时依赖的 Go Telegram Bot API SDK。

`tgbot` 是一个开源 Go SDK，用于构建 Telegram Bot，提供简洁抽象、完整的 Bot API 覆盖、强类型 union 解码、Webhook 支持，以及适合生产环境的长轮询能力。

## 为什么选择 tgbot？

- 完整覆盖 Telegram Bot API 的方法与类型
- 为 Telegram 多态字段和结果提供强类型 union 解码
- 运行时零第三方依赖，仅使用标准库
- 提供文件上传辅助能力，支持 `file_id`、URL、本地路径和 `io.Reader`
- 提供基于 channel 的异步长轮询，支持 `Unsubscribe()` 和可选的非阻塞分发
- 在 `ext` 中提供类似 PTB 的路由层，用于：
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## 快速开始

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

## 安装

```bash
go get github.com/cloudapp3/tgbot
```

## Webhook 路由（`ext`）

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

## 示例

- `examples/quickstart`
- `examples/webhook`
- `examples/commands`
- `examples/async_updates`
- `examples/ext_polling`

## 官方覆盖校验

此 SDK 会对照官方 Telegram Bot API 文档进行校验：`https://core.telegram.org/bots/api`。
截至 `2026-03-06`，该页面上的最新官方版本为 `Bot API 9.5`，发布日期为 `2026-03-01`。

```bash
python3 scripts/verify_telegram_bot_api_coverage.py
```

## 同步 SDK 定义

```bash
go run ./cmd/apigen
```

如果你已经提前下载了官方 HTML 页面：

```bash
go run ./cmd/apigen -html /tmp/telegram_bot_api.html
```

关键文件：

- `sdk_types.go`
- `sdk_methods.go`
- `sdk_unions.go`

## 开发

```bash
go test ./...
go vet ./...
```

## 许可证

MIT
