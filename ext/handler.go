package ext

import (
	"context"
	"regexp"
	"strings"

	tg "github.com/cloudapp3/tgbot"
)

// Context holds routing state for handlers.
type Context struct {
	Bot    *tg.Bot
	Update *Update
}

// UpdateType returns the concrete update type.
func (ctx *Context) UpdateType() UpdateType {
	if ctx == nil || ctx.Update == nil {
		return UpdateTypeUnknown
	}
	return ctx.Update.Type()
}

// EffectiveMessage returns the first message-like payload.
func (ctx *Context) EffectiveMessage() *tg.Message {
	if ctx == nil || ctx.Update == nil {
		return nil
	}
	return ctx.Update.EffectiveMessage()
}

// Command extracts command and args from the effective message text.
func (ctx *Context) Command() (string, string, bool) {
	if ctx == nil || ctx.Update == nil {
		return "", "", false
	}
	return ctx.Update.Command()
}

// Handler is the routing contract used by Application.
type Handler interface {
	Match(*Context) bool
	Handle(context.Context, *Context) error
}

// HandlerFunc is an adapter for handler callbacks.
type HandlerFunc func(context.Context, *Context) error

type handlerAdapter struct {
	match func(*Context) bool
	call  HandlerFunc
}

func (handler *handlerAdapter) Match(ctx *Context) bool {
	if handler == nil || handler.match == nil {
		return false
	}
	return handler.match(ctx)
}

func (handler *handlerAdapter) Handle(ctx context.Context, tgctx *Context) error {
	if handler == nil || handler.call == nil {
		return nil
	}
	return handler.call(ctx, tgctx)
}

// NewAnyHandler registers a handler that receives all updates.
func NewAnyHandler(fn HandlerFunc) Handler {
	return &handlerAdapter{
		match: func(*Context) bool { return true },
		call:  fn,
	}
}

// NewTypeHandler registers a handler for a specific update type.
func NewTypeHandler(updateType UpdateType, fn HandlerFunc) Handler {
	return &handlerAdapter{
		match: func(ctx *Context) bool {
			return ctx != nil && ctx.UpdateType() == updateType
		},
		call: fn,
	}
}

// NewMessageHandler registers a handler for message-like updates.
func NewMessageHandler(filter Filter, fn HandlerFunc) Handler {
	effectiveFilter := filter
	if effectiveFilter == nil {
		effectiveFilter = AnyFilter()
	}

	return &handlerAdapter{
		match: func(ctx *Context) bool {
			if ctx == nil || ctx.EffectiveMessage() == nil {
				return false
			}
			return effectiveFilter.Match(ctx)
		},
		call: fn,
	}
}

// NewCommandHandler registers a handler for a specific command name.
func NewCommandHandler(command string, fn HandlerFunc) Handler {
	expected := normalizeCommand(command)
	return &handlerAdapter{
		match: func(ctx *Context) bool {
			if expected == "" || ctx == nil {
				return false
			}
			commandName, _, ok := ctx.Command()
			if !ok {
				return false
			}
			return normalizeCommand(commandName) == expected
		},
		call: fn,
	}
}

// NewCallbackQueryHandler registers a callback query handler.
// If pattern is nil, every callback query will match.
func NewCallbackQueryHandler(pattern *regexp.Regexp, fn HandlerFunc) Handler {
	return &handlerAdapter{
		match: func(ctx *Context) bool {
			if ctx == nil || ctx.Update == nil || ctx.Update.CallbackQuery == nil {
				return false
			}
			if pattern == nil {
				return true
			}
			return pattern.MatchString(ctx.Update.CallbackQuery.Data)
		},
		call: fn,
	}
}

func normalizeCommand(command string) string {
	normalized := strings.TrimSpace(command)
	normalized = strings.TrimPrefix(normalized, "/")
	if mention := strings.Index(normalized, "@"); mention >= 0 {
		normalized = normalized[:mention]
	}
	return strings.ToLower(strings.TrimSpace(normalized))
}
