# apigen

`apigen` syncs this SDK with the official Telegram Bot API documentation page.

## Usage

From repository root:

```bash
go run ./cmd/apigen
```

If you already downloaded the official HTML page, use it directly:

```bash
go run ./cmd/apigen -html /tmp/telegram_bot_api.html
```

Key files:

- `sdk_types.go`
- `sdk_methods.go`
- `sdk_unions.go`
