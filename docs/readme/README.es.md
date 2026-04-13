# tgbot

[![CI](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/cloudapp3/tgbot/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22%2B-00ADD8?logo=go)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/cloudapp3/tgbot.svg)](https://pkg.go.dev/github.com/cloudapp3/tgbot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/cloudapp3/tgbot/blob/master/LICENSE)

[English](../../README.md) | [简体中文](./README.zh-CN.md) | [Español](./README.es.md) | [日本語](./README.ja.md)

Un SDK ligero y fuertemente tipado de Telegram Bot API para Go, sin dependencias de ejecución de terceros.

`tgbot` es un SDK open source de Go para construir bots de Telegram con abstracciones limpias, cobertura completa de Bot API, decodificación tipada de unions, soporte para webhooks y long polling listo para producción.

## ¿Por qué tgbot?

- Cobertura completa de métodos y tipos de Telegram Bot API
- Decodificación tipada de unions para campos y resultados polimórficos de Telegram
- Cero dependencias de runtime de terceros; solo usa la biblioteca estándar
- Helpers de subida de archivos para `file_id`, URL, ruta local e `io.Reader`
- Long polling asíncrono con suscripciones basadas en channels, `Unsubscribe()` y despacho no bloqueante opcional
- Capa de enrutamiento estilo PTB en `ext` para:
  - `Application`
  - `Handler`
  - `Filter`
  - webhook `http.Handler`

## Inicio Rápido

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

## Instalación

```bash
go get github.com/cloudapp3/tgbot
```

## Enrutamiento Webhook (`ext`)

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

## Ejemplos

- `examples/quickstart`
- `examples/webhook`
- `examples/commands`
- `examples/async_updates`
- `examples/ext_polling`

## Cobertura Oficial

Este SDK se verifica contra la documentación oficial de Telegram Bot API en `https://core.telegram.org/bots/api`.
A fecha de `2026-03-06`, la versión oficial más reciente en esa página es `Bot API 9.5`, publicada el `2026-03-01`.

```bash
python3 scripts/verify_telegram_bot_api_coverage.py
```

## Sincronizar definiciones del SDK

```bash
go run ./cmd/apigen
```

Si ya descargaste la página HTML oficial:

```bash
go run ./cmd/apigen -html /tmp/telegram_bot_api.html
```

Archivos principales:

- `sdk_types.go`
- `sdk_methods.go`
- `sdk_unions.go`

## Desarrollo

```bash
go test ./...
go vet ./...
```

## Licencia

MIT
