# tgbot

[English](../../README.md) | [简体中文](./README.zh-CN.md) | [Español](./README.es.md) | [日本語](./README.ja.md)

一个轻量级的 Go Telegram Bot API SDK，提供完整的更新路由、Webhook 处理支持，以及简洁、适合生产环境的抽象层。

## 特性

- 仓库内已包含对 Telegram Bot API method/type 的完整覆盖，并可对照官方文档校验
- 支持 Telegram 多态字段/结果的强类型 union 解码
- 运行时零第三方依赖，仅使用标准库
- 提供文件上传辅助能力，支持 `file_id`、URL、本地路径和 `io.Reader`
- 基于 `GetUpdates` 提供异步长轮询能力，支持基于 channel 的事件订阅、`Unsubscribe()` 和可选的非阻塞分发
- 在 `ext` 中提供类似 PTB 的路由层：
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## 安装

```bash
go get github.com/cloudapp3/tgbot
```

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
