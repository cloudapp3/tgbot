package ext

import (
	"context"
	"crypto/subtle"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	tg "github.com/cloudapp3/tgbot"
)

const (
	defaultWebhookBodyLimit = int64(4 << 20)
	webhookSecretHeader     = "X-Telegram-Bot-Api-Secret-Token"
)

// ErrorHandler handles errors produced by routed handlers.
type ErrorHandler func(context.Context, *Context, error)

// Option configures Application.
type Option func(*Application)

// Application provides PTB-like update routing on top of tg.Bot.
type Application struct {
	bot             *tg.Bot
	bodyLimit       int64
	continueOnError bool

	mu          sync.RWMutex
	handlers    []Handler
	errorHandle ErrorHandler
}

// NewApplication creates an update dispatcher bound to a bot instance.
func NewApplication(bot *tg.Bot, opts ...Option) (*Application, error) {
	if bot == nil {
		return nil, fmt.Errorf("telegram bot is nil")
	}

	app := &Application{
		bot:             bot,
		bodyLimit:       defaultWebhookBodyLimit,
		continueOnError: true,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(app)
		}
	}
	if app.bodyLimit <= 0 {
		app.bodyLimit = defaultWebhookBodyLimit
	}
	return app, nil
}

// WithWebhookBodyLimit overrides max webhook body size.
func WithWebhookBodyLimit(limit int64) Option {
	return func(app *Application) {
		if app != nil && limit > 0 {
			app.bodyLimit = limit
		}
	}
}

// WithErrorHandler sets a global error callback.
func WithErrorHandler(handler ErrorHandler) Option {
	return func(app *Application) {
		if app != nil {
			app.errorHandle = handler
		}
	}
}

// WithContinueOnError controls whether dispatcher continues after a handler error.
func WithContinueOnError(enabled bool) Option {
	return func(app *Application) {
		if app != nil {
			app.continueOnError = enabled
		}
	}
}

// Bot returns the bound bot instance.
func (app *Application) Bot() *tg.Bot {
	if app == nil {
		return nil
	}
	return app.bot
}

// AddHandler appends a handler to the routing chain.
func (app *Application) AddHandler(handler Handler) {
	if app == nil || handler == nil {
		return
	}
	app.mu.Lock()
	defer app.mu.Unlock()
	app.handlers = append(app.handlers, handler)
}

// SetErrorHandler updates the global handler error callback.
func (app *Application) SetErrorHandler(handler ErrorHandler) {
	if app == nil {
		return
	}
	app.mu.Lock()
	defer app.mu.Unlock()
	app.errorHandle = handler
}

// ProcessUpdate routes an update through all matching handlers.
func (app *Application) ProcessUpdate(ctx context.Context, update *Update) error {
	if app == nil {
		return fmt.Errorf("application is nil")
	}
	if update == nil {
		return fmt.Errorf("update is nil")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	contextValue := &Context{
		Bot:    app.bot,
		Update: update,
	}
	handlers, errorHandler, continueOnError := app.snapshotRouting()

	var firstErr error
	for _, handler := range handlers {
		if handler == nil || !handler.Match(contextValue) {
			continue
		}

		if err := handler.Handle(ctx, contextValue); err != nil {
			if firstErr == nil {
				firstErr = err
			}
			if errorHandler != nil {
				errorHandler(ctx, contextValue, err)
			}
			if !continueOnError {
				return err
			}
		}
	}
	return firstErr
}

// WebhookHandler returns an http.Handler for Telegram webhook callbacks.
// If secretToken is non-empty, requests must pass X-Telegram-Bot-Api-Secret-Token.
func (app *Application) WebhookHandler(secretToken string) http.Handler {
	expectedSecret := strings.TrimSpace(secretToken)

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.Header().Set("Allow", http.MethodPost)
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if expectedSecret != "" {
			actual := strings.TrimSpace(request.Header.Get(webhookSecretHeader))
			if subtle.ConstantTimeCompare([]byte(actual), []byte(expectedSecret)) != 1 {
				http.Error(writer, "unauthorized", http.StatusUnauthorized)
				return
			}
		}

		update, err := decodeUpdateFromRequest(request, app.effectiveBodyLimit())
		if err != nil {
			http.Error(writer, "bad request", http.StatusBadRequest)
			return
		}
		if err := app.ProcessUpdate(request.Context(), update); err != nil {
			http.Error(writer, "handler error", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("ok"))
	})
}

func (app *Application) snapshotRouting() ([]Handler, ErrorHandler, bool) {
	app.mu.RLock()
	defer app.mu.RUnlock()

	handlers := make([]Handler, len(app.handlers))
	copy(handlers, app.handlers)
	return handlers, app.errorHandle, app.continueOnError
}

func (app *Application) effectiveBodyLimit() int64 {
	if app == nil || app.bodyLimit <= 0 {
		return defaultWebhookBodyLimit
	}
	return app.bodyLimit
}

func decodeUpdateFromRequest(request *http.Request, limit int64) (*Update, error) {
	if request == nil || request.Body == nil {
		return nil, fmt.Errorf("request body is nil")
	}
	defer request.Body.Close()

	body, err := io.ReadAll(io.LimitReader(request.Body, limit+1))
	if err != nil {
		return nil, fmt.Errorf("read webhook body: %w", err)
	}
	if int64(len(body)) > limit {
		return nil, fmt.Errorf("webhook body exceeds limit")
	}
	return DecodeUpdate(body)
}
