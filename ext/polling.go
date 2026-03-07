package ext

import (
	"context"
	"fmt"

	tg "github.com/cloudapp3/tgbot"
)

// PollingOption configures Application.RunPolling.
type PollingOption = tg.UpdatePollerOption

// WithPollingAllowedUpdates converts ext update types into poller allowed_updates values.
func WithPollingAllowedUpdates(updateTypes ...UpdateType) PollingOption {
	rootTypes := make([]tg.UpdateType, 0, len(updateTypes))
	for _, updateType := range updateTypes {
		if updateType == UpdateTypeUnknown {
			continue
		}
		rootTypes = append(rootTypes, tg.UpdateType(updateType))
	}
	return tg.WithPollerAllowedUpdates(rootTypes...)
}

// WithPollingNonBlockingDispatch enables non-blocking dispatch for Application.RunPolling.
func WithPollingNonBlockingDispatch() PollingOption {
	return tg.WithPollerNonBlockingDispatch()
}

// RunPolling starts long polling in the background and routes updates through the application.
func (app *Application) RunPolling(ctx context.Context, opts ...PollingOption) error {
	if app == nil {
		return fmt.Errorf("application is nil")
	}
	if app.bot == nil {
		return fmt.Errorf("telegram bot is nil")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	poller, err := app.bot.StartUpdatePoller(ctx, opts...)
	if err != nil {
		return err
	}
	defer poller.Stop()

	updatesCh := poller.Updates()
	errorsCh := poller.Errors()
	for updatesCh != nil || errorsCh != nil {
		select {
		case <-ctx.Done():
			return nil
		case err, ok := <-errorsCh:
			if !ok {
				errorsCh = nil
				continue
			}
			if err == nil {
				continue
			}
			app.reportError(ctx, nil, err)
			if !app.shouldContinueOnError() {
				return err
			}
		case update, ok := <-updatesCh:
			if !ok {
				updatesCh = nil
				continue
			}
			if err := app.ProcessUpdate(ctx, WrapUpdate(update)); err != nil && !app.shouldContinueOnError() {
				return err
			}
		}
	}
	return nil
}

func (app *Application) shouldContinueOnError() bool {
	if app == nil {
		return false
	}
	_, _, continueOnError := app.snapshotRouting()
	return continueOnError
}

func (app *Application) reportError(ctx context.Context, update *Update, err error) {
	if app == nil || err == nil {
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}
	_, errorHandler, _ := app.snapshotRouting()
	if errorHandler == nil {
		return
	}
	errorHandler(ctx, &Context{Bot: app.bot, Update: update}, err)
}
