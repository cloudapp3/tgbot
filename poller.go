package tgbot

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	defaultUpdatePollerTimeout     = int64(25)
	defaultUpdatePollerLimit       = int64(100)
	defaultUpdatePollerBuffer      = 64
	defaultUpdatePollerErrorBuffer = 16
	defaultUpdatePollerRetryDelay  = 2 * time.Second
)

// UpdateFilter matches updates for subscription fan-out.
type UpdateFilter func(Update) bool

// UpdatePollerOption configures an UpdatePoller.
type UpdatePollerOption func(*UpdatePoller)

// UpdateDispatchError reports that a non-blocking poller dropped an update.
type UpdateDispatchError struct {
	UpdateID     int64
	UpdateType   UpdateType
	Target       string
	SubscriberID uint64
}

func (err *UpdateDispatchError) Error() string {
	if err == nil {
		return "telegram update dispatch error"
	}
	if err.SubscriberID > 0 {
		return fmt.Sprintf("telegram update dropped: target=%s subscriber_id=%d update_id=%d type=%s", err.Target, err.SubscriberID, err.UpdateID, err.UpdateType)
	}
	return fmt.Sprintf("telegram update dropped: target=%s update_id=%d type=%s", err.Target, err.UpdateID, err.UpdateType)
}

// UpdateSubscription is a cancellable subscription to poller updates.
type UpdateSubscription struct {
	poller *UpdatePoller
	sub    *subscriber
	once   sync.Once
}

// ID returns the internal subscription identifier.
func (subscription *UpdateSubscription) ID() uint64 {
	if subscription == nil || subscription.sub == nil {
		return 0
	}
	return subscription.sub.id
}

// Updates returns the subscription channel.
func (subscription *UpdateSubscription) Updates() <-chan Update {
	if subscription == nil || subscription.sub == nil {
		return nil
	}
	return subscription.sub.ch
}

// Unsubscribe removes the subscription and closes its channel.
func (subscription *UpdateSubscription) Unsubscribe() {
	if subscription == nil {
		return
	}
	subscription.once.Do(func() {
		if subscription.poller == nil || subscription.sub == nil {
			return
		}
		subscription.poller.unsubscribe(subscription.sub)
	})
}

// UpdatePoller runs getUpdates in the background and fans updates out to channels.
type UpdatePoller struct {
	bot *Bot

	params              GetUpdatesParams
	updateBuffer        int
	errorBuffer         int
	retryDelay          time.Duration
	nonBlockingDispatch bool

	updates chan Update
	errors  chan error
	done    chan struct{}

	mu               sync.RWMutex
	subscribers      []*subscriber
	nextSubscriberID uint64
	closed           bool

	startOnce sync.Once
	cancel    context.CancelFunc
	wg        sync.WaitGroup
}

type subscriber struct {
	id     uint64
	filter UpdateFilter
	ch     chan Update

	mu     sync.RWMutex
	closed bool
}

// WithPollerParams overrides the base getUpdates parameters used by the poller.
func WithPollerParams(params GetUpdatesParams) UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller == nil {
			return
		}
		poller.params = cloneGetUpdatesParams(params)
	}
}

// WithPollerAllowedUpdates sets allowed_updates using typed update names.
func WithPollerAllowedUpdates(updateTypes ...UpdateType) UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller == nil {
			return
		}
		allowed := make([]string, 0, len(updateTypes))
		for _, updateType := range updateTypes {
			if updateType == UpdateTypeUnknown {
				continue
			}
			allowed = append(allowed, string(updateType))
		}
		poller.params.AllowedUpdates = allowed
	}
}

// WithPollerBuffer sets the buffer size for the all-updates channel.
func WithPollerBuffer(size int) UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller == nil || size < 0 {
			return
		}
		poller.updateBuffer = size
	}
}

// WithPollerErrorBuffer sets the buffer size for the error channel.
func WithPollerErrorBuffer(size int) UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller == nil || size < 0 {
			return
		}
		poller.errorBuffer = size
	}
}

// WithPollerRetryDelay sets the delay before retrying after getUpdates errors.
func WithPollerRetryDelay(delay time.Duration) UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller == nil || delay < 0 {
			return
		}
		poller.retryDelay = delay
	}
}

// WithPollerNonBlockingDispatch enables non-blocking fan-out. Slow consumers drop updates instead of blocking the poller.
func WithPollerNonBlockingDispatch() UpdatePollerOption {
	return func(poller *UpdatePoller) {
		if poller != nil {
			poller.nonBlockingDispatch = true
		}
	}
}

// NewUpdatePoller creates a long-polling helper on top of Bot.GetUpdates.
func NewUpdatePoller(bot *Bot, opts ...UpdatePollerOption) (*UpdatePoller, error) {
	if bot == nil {
		return nil, fmt.Errorf("telegram bot is nil")
	}

	poller := &UpdatePoller{
		bot:          bot,
		params:       GetUpdatesParams{Timeout: defaultUpdatePollerTimeout, Limit: defaultUpdatePollerLimit},
		updateBuffer: defaultUpdatePollerBuffer,
		errorBuffer:  defaultUpdatePollerErrorBuffer,
		retryDelay:   defaultUpdatePollerRetryDelay,
		done:         make(chan struct{}),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(poller)
		}
	}
	if poller.params.Timeout <= 0 {
		poller.params.Timeout = defaultUpdatePollerTimeout
	}
	if poller.params.Limit <= 0 {
		poller.params.Limit = defaultUpdatePollerLimit
	}
	poller.updates = make(chan Update, poller.updateBuffer)
	poller.errors = make(chan error, poller.errorBuffer)
	return poller, nil
}

// StartUpdatePoller creates and starts an UpdatePoller.
func (bot *Bot) StartUpdatePoller(ctx context.Context, opts ...UpdatePollerOption) (*UpdatePoller, error) {
	poller, err := NewUpdatePoller(bot, opts...)
	if err != nil {
		return nil, err
	}
	poller.Start(ctx)
	return poller, nil
}

// Start launches the background polling goroutine. Calling Start more than once is a no-op.
func (poller *UpdatePoller) Start(ctx context.Context) {
	if poller == nil {
		return
	}
	poller.startOnce.Do(func() {
		if ctx == nil {
			ctx = context.Background()
		}
		runCtx, cancel := context.WithCancel(ctx)
		poller.cancel = cancel
		poller.wg.Add(1)
		go func() {
			defer poller.wg.Done()
			defer poller.shutdown()
			poller.loop(runCtx)
		}()
	})
}

// Stop cancels the background poller and waits for shutdown.
func (poller *UpdatePoller) Stop() {
	if poller == nil {
		return
	}
	if poller.cancel != nil {
		poller.cancel()
	}
	poller.wg.Wait()
}

// Done is closed when the poller exits.
func (poller *UpdatePoller) Done() <-chan struct{} {
	if poller == nil {
		closed := make(chan struct{})
		close(closed)
		return closed
	}
	return poller.done
}

// Updates returns the fan-out channel containing every update.
func (poller *UpdatePoller) Updates() <-chan Update {
	if poller == nil {
		return nil
	}
	return poller.updates
}

// Errors returns polling errors. Errors are dropped when the channel buffer is full.
func (poller *UpdatePoller) Errors() <-chan error {
	if poller == nil {
		return nil
	}
	return poller.errors
}

// Subscribe registers a filtered update channel.
func (poller *UpdatePoller) Subscribe(buffer int, filter UpdateFilter) <-chan Update {
	subscription := poller.SubscribeHandle(buffer, filter)
	if subscription == nil {
		return nil
	}
	return subscription.Updates()
}

// SubscribeHandle registers a filtered update subscription with Unsubscribe support.
func (poller *UpdatePoller) SubscribeHandle(buffer int, filter UpdateFilter) *UpdateSubscription {
	if poller == nil {
		return nil
	}
	if buffer < 0 {
		buffer = 0
	}
	if filter == nil {
		filter = func(Update) bool { return true }
	}

	sub := &subscriber{
		filter: filter,
		ch:     make(chan Update, buffer),
	}

	poller.mu.Lock()
	defer poller.mu.Unlock()
	if poller.closed {
		sub.close()
		return &UpdateSubscription{poller: poller, sub: sub}
	}
	poller.nextSubscriberID++
	sub.id = poller.nextSubscriberID
	poller.subscribers = append(poller.subscribers, sub)
	return &UpdateSubscription{poller: poller, sub: sub}
}

// SubscribeTypes registers a channel that receives only the requested update types.
func (poller *UpdatePoller) SubscribeTypes(buffer int, updateTypes ...UpdateType) <-chan Update {
	subscription := poller.SubscribeTypesHandle(buffer, updateTypes...)
	if subscription == nil {
		return nil
	}
	return subscription.Updates()
}

// SubscribeTypesHandle registers a typed update subscription with Unsubscribe support.
func (poller *UpdatePoller) SubscribeTypesHandle(buffer int, updateTypes ...UpdateType) *UpdateSubscription {
	set := make(map[UpdateType]struct{}, len(updateTypes))
	for _, updateType := range updateTypes {
		if updateType == UpdateTypeUnknown {
			continue
		}
		set[updateType] = struct{}{}
	}
	return poller.SubscribeHandle(buffer, func(update Update) bool {
		_, ok := set[update.Type()]
		return ok
	})
}

func (poller *UpdatePoller) loop(ctx context.Context) {
	params := cloneGetUpdatesParams(poller.params)
	for {
		updates, err := poller.bot.GetUpdates(ctx, &params)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			poller.emitError(err)
			if !sleepContext(ctx, poller.retryDelay) {
				return
			}
			continue
		}

		for _, update := range updates {
			if update.UpdateID >= params.Offset {
				params.Offset = update.UpdateID + 1
			}
			if !poller.dispatch(ctx, update) {
				return
			}
		}
	}
}

func (poller *UpdatePoller) dispatch(ctx context.Context, update Update) bool {
	if !poller.sendUpdate(ctx, poller.updates, update, "updates", 0) {
		return false
	}
	for _, sub := range poller.matchingSubscribers(update) {
		if !sub.send(ctx, poller, update) {
			return false
		}
	}
	return true
}

func (poller *UpdatePoller) sendUpdate(ctx context.Context, ch chan Update, update Update, target string, subscriberID uint64) bool {
	if poller == nil {
		return false
	}
	if poller.nonBlockingDispatch {
		select {
		case ch <- update:
			return true
		default:
			if ctx.Err() != nil {
				return false
			}
			poller.emitError(&UpdateDispatchError{
				UpdateID:     update.UpdateID,
				UpdateType:   update.Type(),
				Target:       target,
				SubscriberID: subscriberID,
			})
			return true
		}
	}

	select {
	case ch <- update:
		return true
	case <-ctx.Done():
		return false
	}
}

func (poller *UpdatePoller) matchingSubscribers(update Update) []*subscriber {
	poller.mu.RLock()
	defer poller.mu.RUnlock()

	subs := make([]*subscriber, 0, len(poller.subscribers))
	for _, sub := range poller.subscribers {
		if sub == nil {
			continue
		}
		if sub.filter == nil || sub.filter(update) {
			subs = append(subs, sub)
		}
	}
	return subs
}

func (poller *UpdatePoller) unsubscribe(target *subscriber) {
	if poller == nil || target == nil {
		return
	}

	poller.mu.Lock()
	for index, sub := range poller.subscribers {
		if sub != target {
			continue
		}
		poller.subscribers = append(poller.subscribers[:index], poller.subscribers[index+1:]...)
		break
	}
	poller.mu.Unlock()
	target.close()
}

func (poller *UpdatePoller) emitError(err error) {
	if poller == nil || err == nil {
		return
	}
	select {
	case poller.errors <- err:
	default:
	}
}

func (poller *UpdatePoller) shutdown() {
	poller.mu.Lock()
	if poller.closed {
		poller.mu.Unlock()
		return
	}
	poller.closed = true
	subscribers := append([]*subscriber(nil), poller.subscribers...)
	poller.subscribers = nil
	updates := poller.updates
	errors := poller.errors
	done := poller.done
	poller.mu.Unlock()

	close(updates)
	for _, sub := range subscribers {
		if sub != nil {
			sub.close()
		}
	}
	close(errors)
	close(done)
}

func (sub *subscriber) send(ctx context.Context, poller *UpdatePoller, update Update) bool {
	if sub == nil || poller == nil {
		return true
	}
	sub.mu.RLock()
	defer sub.mu.RUnlock()
	if sub.closed {
		return true
	}
	return poller.sendUpdate(ctx, sub.ch, update, "subscription", sub.id)
}

func (sub *subscriber) close() {
	if sub == nil {
		return
	}
	sub.mu.Lock()
	defer sub.mu.Unlock()
	if sub.closed {
		return
	}
	sub.closed = true
	close(sub.ch)
}

func cloneGetUpdatesParams(params GetUpdatesParams) GetUpdatesParams {
	cloned := params
	if params.AllowedUpdates != nil {
		cloned.AllowedUpdates = append([]string(nil), params.AllowedUpdates...)
	}
	return cloned
}

func sleepContext(ctx context.Context, delay time.Duration) bool {
	if delay <= 0 {
		select {
		case <-ctx.Done():
			return false
		default:
			return true
		}
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}
