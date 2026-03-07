# tgbot

[English](../../README.md) | [简体中文](./README.zh-CN.md) | [Español](./README.es.md) | [日本語](./README.ja.md)

`tgbot` は、完全な update ルーティング、Webhook ハンドラー対応、本番利用しやすいクリーンな抽象化を備えた軽量な Go 向け Telegram Bot API SDK です。

## 特徴

- このリポジトリに Telegram Bot API の method/type を網羅し、公式ドキュメントとの照合も可能
- Telegram の多態フィールド・結果に対する型付き union デコードをサポート
- 実行時のサードパーティ依存なし、標準ライブラリのみ使用
- ファイルアップロード用ヘルパー（`file_id`、URL、ローカルパス、`io.Reader`）を提供
- `GetUpdates` の上に channel ベースの購読、`Unsubscribe()`、任意の非ブロッキング分配を備えた非同期ロングポーリングヘルパーを提供
- `ext` に PTB 風のルーティング層を提供:
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## インストール

```bash
go get github.com/cloudapp3/tgbot
```

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
