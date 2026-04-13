# tgbot

[![CI](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22%2B-00ADD8?logo=go)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/cloudapp3/tgbot.svg)](https://pkg.go.dev/github.com/cloudapp3/tgbot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/cloudapp3/tgbot/blob/master/LICENSE)

[English](../../README.md) | [简体中文](./README.zh-CN.md) | [Español](./README.es.md) | [日本語](./README.ja.md)

軽量・型安全・ランタイム外部依存なしの Go 向け Telegram Bot API SDK。

`tgbot` は、クリーンな抽象化、完全な Bot API カバレッジ、型安全な union デコード、Webhook 対応、そして本番運用しやすいロングポーリングを備えた、Telegram Bot 開発向けのオープンソース Go SDK です。

## なぜ tgbot なのか？

- Telegram Bot API のメソッドと型を幅広くカバー
- Telegram の多態フィールドと結果に対する型安全な union デコードをサポート
- 実行時のサードパーティ依存なし、標準ライブラリのみ使用
- ファイルアップロード用ヘルパー（`file_id`、URL、ローカルパス、`io.Reader`）を提供
- channel ベースの購読、`Unsubscribe()`、任意の非ブロッキング分配を備えた非同期ロングポーリングを提供
- `ext` に PTB 風のルーティング層を提供し、以下を扱えます:
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## クイックスタート

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

## インストール

```bash
go get github.com/cloudapp3/tgbot
```

## Webhook ルーティング（`ext`）

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

## サンプル

- `examples/quickstart`
- `examples/webhook`
- `examples/commands`
- `examples/async_updates`
- `examples/ext_polling`

## 公式カバレッジ

この SDK は公式 Telegram Bot API ドキュメント `https://core.telegram.org/bots/api` を基準に検証できます。
`2026-03-06` 時点で、そのページの最新公式リリースは `Bot API 9.5` で、公開日は `2026-03-01` です。

```bash
python3 scripts/verify_telegram_bot_api_coverage.py
```

## SDK 定義の同期

```bash
go run ./cmd/apigen
```

すでに公式 HTML ページをダウンロード済みの場合:

```bash
go run ./cmd/apigen -html /tmp/telegram_bot_api.html
```

主要ファイル:

- `sdk_types.go`
- `sdk_methods.go`
- `sdk_unions.go`

## 開発

```bash
go test ./...
go vet ./...
```

## ライセンス

MIT
